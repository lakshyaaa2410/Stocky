package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lakshyaaa2410/stocky/initializers"
	"github.com/lakshyaaa2410/stocky/models"
	"github.com/lakshyaaa2410/stocky/utilities"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetStockRewardsToday(ginCtx *gin.Context) {

	// Converting UserID From String To Int
	userIdStr := ginCtx.Param("userId")
	userId, err := strconv.Atoi(userIdStr)

	// Checking If Any Error While Conversion.
	if err != nil {
		logrus.Error("Error Pasring User ID")
		return
	}

	// Variable To Store Reward
	var reward models.Reward

	// Checking If User With That ID Exists
	err = initializers.DB.Where("user_id = ?", userId).First(&reward).Error
	if err != nil {
		ginCtx.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "No User Found With User ID: " + userIdStr,
		})
		return
	}

	// Getting Formatted Date (yyyy-mm-dd) Formatt.
	_, today := utilities.GetDayAndTime()

	// Variable To Store All The Rewards.
	var rewards []models.Reward

	// Querying Database To Get All The Reward Of User, For Today.
	response := initializers.DB.Where("user_id = ? AND DATE(rewarded_at) = ?", reward.UserID, today).Find(&rewards)

	// Checking If Some Database Errors.
	if response.Error != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": response.Error.Error(),
		})

		return
	}

	// Sending Back All The Rewards Of Today, For A User.
	ginCtx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"count":   len(rewards),
		"rewards": rewards,
	})

}

func AddReward(ginCtx *gin.Context) {

	// Initializing A New Reward Variable
	var reward models.Reward

	// Binding The Incoming JSON Body Into Reward Variable
	err := ginCtx.ShouldBindJSON(&reward)

	// Checking If There Is Any Error In The Request Body.
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	// Reward Model Variable To Store Any Duplicate Entry For "Onboarding"
	var existingReward models.Reward

	// Making Sure "Action" And "Stock Symbol" Are Always First Letter Capitalized
	reward.Action = utilities.UpperCaseFirstLetter(reward.Action)
	reward.StockSymbol = utilities.UpperCaseFirstLetter(reward.StockSymbol)

	// 3. Checking If The User Has Already Onboarded Or Not
	if action := reward.Action; action == "Onboarding" {
		fmt.Println("In Onboarding Condition")

		err := initializers.DB.Where("user_id = ? AND action = ?", reward.UserID, "Onboarding").First(&existingReward).Error

		// Means A Result With Same User And "Onboarding" Status Was Found
		if err == nil {
			ginCtx.JSON(http.StatusConflict, gin.H{
				"status":  "failed",
				"message": "Onboarding Reward Already Given",
			})
			return
		}
	}

	// Setting Up The Curent IST Time
	reward.RewardedAt, _ = utilities.GetDayAndTime()

	// Using Transactions To Safely Create Reward + Ledger Entry.
	err = initializers.DB.Transaction(func(tx *gorm.DB) error {

		// 1. Creating A Reward Entry First
		response := tx.Create(&reward)

		// Checking For Possible Errors After Creation Of Reward.
		if response.Error != nil {
			ginCtx.JSON(http.StatusConflict, gin.H{
				"status":  "failed",
				"message": response.Error.Error(),
			})

			return err
		}

		// 2.A Initializing A New Stock Price Struct Variable (To Store Stock Price)
		var stockPrice models.StockPrice

		// Fetching The Price Of A Stock, Using Stock Symbol.
		err = tx.Where("stock_symbol = ?", reward.StockSymbol).First(&stockPrice).Error

		// Checking For Possible Errors, If Any, We Rollback.
		if err != nil {
			return err
		}

		// Fetching All The Ledger Entry Queries, Using A Helper Function.
		var ledgerQueries = createLedgerEntries(&reward, &stockPrice)

		// Traversing Through All The Queries, And Creating A New Entry For Each Type.
		for _, query := range ledgerQueries {
			err := tx.Create(&query).Error
			if err != nil {
				return err
			}
		}

		// Commiting A Transaction
		return nil

	})

	// Checking Possible Rollback
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	// Returning JSON Response, With
	ginCtx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": gin.H{
			"reward": reward,
		},
	})
}

func createLedgerEntries(reward *models.Reward, stockPrice *models.StockPrice) []models.Ledger {
	// Helper Method To Calculate The Deductions Like GST And Brokerage
	brokerage, gst := getDeductions(reward.Shares, stockPrice.StockPrice)

	// Creating 8 Entries For The Ledger Table For (Stock, Cash, GST, And Brokerage)(4 - Credit, 4- Debit)
	var ledgerEntries = []models.Ledger{
		// These 2 Entries Are For Stocks
		{
			RewardID:        reward.ID,
			UserID:          reward.UserID,
			TransactionType: "Stocks",
			Amount:          reward.Shares,
			AmountUnit:      "Stock",
			FlowType:        "Debit",
			Account:         "Company",
			Action:          reward.Action,
		},
		{
			RewardID:        reward.ID,
			UserID:          reward.UserID,
			TransactionType: "Stocks",
			Amount:          reward.Shares,
			AmountUnit:      "Stock",
			FlowType:        "Credit",
			Account:         "User",
			Action:          reward.Action,
		},

		// These 2 Entries Are For Cash Transaction
		{
			RewardID:        reward.ID,
			UserID:          reward.UserID,
			TransactionType: "Cash",
			Amount:          reward.Shares * stockPrice.StockPrice,
			AmountUnit:      "INR",
			FlowType:        "Credit",
			Account:         "Company",
			Action:          reward.Action,
		},
		{
			RewardID:        reward.ID,
			UserID:          reward.UserID,
			TransactionType: "Cash",
			Amount:          reward.Shares * stockPrice.StockPrice,
			AmountUnit:      "INR",
			FlowType:        "Debit",
			Account:         "User",
			Action:          reward.Action,
		},

		// These 2 Entries Are For GST
		{
			RewardID:        reward.ID,
			UserID:          reward.UserID,
			TransactionType: "GST",
			Amount:          gst,
			AmountUnit:      "INR",
			FlowType:        "Debit",
			Account:         "Company",
			Action:          reward.Action,
		},
		{
			RewardID:        reward.ID,
			UserID:          reward.UserID,
			TransactionType: "GST",
			Amount:          gst,
			AmountUnit:      "INR",
			FlowType:        "Credit",
			Account:         "Government",
			Action:          reward.Action,
		},

		// These 2 Entries Are For Brokerage
		{
			RewardID:        reward.ID,
			UserID:          reward.UserID,
			TransactionType: "Brokerage",
			Amount:          brokerage,
			AmountUnit:      "INR",
			FlowType:        "Debit",
			Account:         "Company",
			Action:          reward.Action,
		},
		{
			RewardID:        reward.ID,
			UserID:          reward.UserID,
			TransactionType: "Brokerage",
			Amount:          brokerage,
			AmountUnit:      "INR",
			FlowType:        "Credit",
			Account:         "Broker",
			Action:          reward.Action,
		},
	}

	return ledgerEntries
}

func getDeductions(quantity float64, price float64) (float64, float64) {

	// Calculation Of Brokerage (Assuming 5% Deduction)
	var brokerageValue = (quantity * price) * 0.05

	// Calculation Of GST (Assuming 18% Deduction)
	var gstValue = brokerageValue * 0.18

	return brokerageValue, gstValue
}
