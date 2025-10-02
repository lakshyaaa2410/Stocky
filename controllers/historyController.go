package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lakshyaaa2410/stocky/initializers"
	"github.com/sirupsen/logrus"
)

// Struct To Store The Date And Value Of User's Past Rewards
type UserHistory struct {
	Date  string  `json:"date" gorm:"column:date"`
	Value float64 `json:"value" gorm:"column:value"`
}

func GetStockHistory(ginCtx *gin.Context) {
	// Converting UserID From String To Int
	userIdStr := ginCtx.Param("userId")
	userId, err := strconv.Atoi(userIdStr)

	// Checking For Potential Errors, After Conversion.
	if err != nil {
		logrus.Error("Error Parsing User Id")
		return
	}

	// Variable To Store User History
	var userHistory []UserHistory

	// Helper Method To Get User's Stock History Based On UserId.
	err = getUserStockHistory(userId, &userHistory)

	// Checking For Potential Errors After Data Fetching.
	if err != nil {
		logrus.Error("Error Fetching User History")
		return
	}

	// If None, Proceeding By Sending Response.
	ginCtx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   userId,
		"data":   userHistory,
	})
}

func getUserStockHistory(userId int, userHistory *[]UserHistory) error {
	// Creating A Raw SQL Query To Fetch User's Stock History.
	query := `SELECT DATE(rewardTable.rewarded_at) AS date,
            ROUND(SUM(rewardTable.shares * stockPriceTable.stock_price), 2) AS value
            FROM rewards rewardTable
            JOIN stock_prices stockPriceTable 
            ON rewardTable.stock_symbol = stockPriceTable.stock_symbol
            WHERE rewardTable.user_id = ?
            AND DATE(rewardTable.rewarded_at) < CURRENT_DATE
            GROUP BY DATE(rewardTable.rewarded_at)
            ORDER BY DATE(rewardTable.rewarded_at);`

	// Executing The Raw Query And Saving The Result Into userHistory Variable
	// Returning The Error (If Any)
	return initializers.DB.Raw(query, userId).Scan(userHistory).Error
}
