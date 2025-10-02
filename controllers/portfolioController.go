package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lakshyaaa2410/stocky/initializers"
)

type UserPortfolio struct {
	StockSymbol string  `json:"stockSymbol" gorm:"stock_symbol"`
	Shares      float64 `json:"shares" gorm:"shares"`
	Value       float64 `json:"value" gorm:"value"`
}

func GetUserPortfolio(ginCtx *gin.Context) {

	userIdStr := ginCtx.Param("userId")
	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Error Parsing User ID",
		})
		return
	}

	var userPortfolio []UserPortfolio
	err = getUserPortfolioValue(userId, &userPortfolio)

	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Error Fetching User Portfolio",
		})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"count":  len(userPortfolio),
		"data":   userPortfolio,
	})
}

func getUserPortfolioValue(userId int, userPortfolio *[]UserPortfolio) error {

	query := `SELECT rewardTable.stock_symbol,
			SUM(rewardTable.shares) AS shares,
			ROUND(SUM(rewardTable.shares * stockPriceTable.stock_price), 2) AS value
			FROM rewards rewardTable
			JOIN stock_prices stockPriceTable 
			ON rewardTable.stock_symbol = stockPriceTable.stock_symbol
			WHERE rewardTable.user_id = ?
			GROUP BY rewardTable.stock_symbol`

	return initializers.DB.Raw(query, userId).Scan(&userPortfolio).Error
}
