package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lakshyaaa2410/stocky/initializers"
)

// Struct To Store The Stock Symbol, Shares, And Value Of User's Portfolio.
type UserPortfolio struct {
	StockSymbol string  `json:"stockSymbol" gorm:"stock_symbol"`
	Shares      float64 `json:"shares" gorm:"shares"`
	Value       float64 `json:"value" gorm:"value"`
}

func GetUserPortfolio(ginCtx *gin.Context) {

	// Converting UserID From String To Int.
	userIdStr := ginCtx.Param("userId")
	userId, err := strconv.Atoi(userIdStr)

	// Checking For Potential Errors During Conversion.
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Error Parsing User ID",
		})
		return
	}

	// Variable To Store User Portfolio.
	var userPortfolio []UserPortfolio

	// Helper Method To Get User's Portfolio Based On userId.
	err = getUserPortfolioValue(userId, &userPortfolio)

	// Checking For Potential Errors After Data Fetchnig.
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Error Fetching User Portfolio",
		})
		return
	}

	// Sending Back Response, If No Errors.
	ginCtx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"count":  len(userPortfolio),
		"data":   userPortfolio,
	})
}

func getUserPortfolioValue(userId int, userPortfolio *[]UserPortfolio) error {
	// Creating A Raw SQL Query To Fetch User's Portfolio
	query := `SELECT rewardTable.stock_symbol,
			SUM(rewardTable.shares) AS shares,
			ROUND(SUM(rewardTable.shares * stockPriceTable.stock_price), 2) AS value
			FROM rewards rewardTable
			JOIN stock_prices stockPriceTable 
			ON rewardTable.stock_symbol = stockPriceTable.stock_symbol
			WHERE rewardTable.user_id = ?
			GROUP BY rewardTable.stock_symbol`

	// Executing The Raw Query And Saving The Result Into userPortfolio Variable
	// Returning The Error (If Any)
	return initializers.DB.Raw(query, userId).Scan(&userPortfolio).Error
}
