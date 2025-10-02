package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lakshyaaa2410/stocky/initializers"
	"github.com/sirupsen/logrus"
)

type UserHistory struct {
	Date  time.Time `json:"date" gorm:"date"`
	Value float64   `json:"value" gorm:"value"`
}

func GetStockHistory(ginCtx *gin.Context) {
	userIdStr := ginCtx.Param("userId")
	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		logrus.Error("Error Parsing User Id")
		return
	}

	var UserHistory []UserHistory
	err = getUserStockHistory(userId, &UserHistory)

	if err != nil {
		logrus.Error("Error Fetching User History")
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   userId,
		"data":   UserHistory,
	})
}

func getUserStockHistory(userId int, userHistory *[]UserHistory) error {

	query := `SELECT DATE(rewardTable.rewarded_at) AS date,
    		SUM(rewardTable.shares * stockPriceTable.stock_price) AS value
			FROM rewards rewardTable
			JOIN stock_prices stockPriceTable ON rewardTable.stock_symbol = stockPriceTable.stock_symbol
			WHERE rewardTable.user_id = ?
  			AND DATE(rewardTable.rewarded_at) < CURRENT_DATE
			GROUP BY DATE(rewardTable.rewarded_at)
			ORDER BY DATE(rewardTable.rewarded_at)`

	return initializers.DB.Raw(query, userId).Scan(&userHistory).Error
}
