package controllers

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lakshyaaa2410/stocky/initializers"
	"github.com/lakshyaaa2410/stocky/models"
	"github.com/sirupsen/logrus"
)

func UpdateStockPrices(ginCtx *gin.Context) {

	var stockPrices []models.StockPrice

	err := initializers.DB.Find(&stockPrices).Error
	if err != nil {
		logrus.Error("Error Fetching Stock Prices")
		return
	}

	for _, stockPrices := range stockPrices {
		currentStockPrice := stockPrices.StockPrice
		newStockPrice := getRandomStockPrice(currentStockPrice, 10)

		err := initializers.DB.Model(&stockPrices).Update("StockPrice", newStockPrice).Error

		if err != nil {
			logrus.Error("Error Updating Price")
			continue
		}

		fmt.Println("The Current And New Stock Prices Are", currentStockPrice, newStockPrice)
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Stock Prices Updated",
	})

}

func getRandomStockPrice(currStockPrice float64, diffBy int) float64 {

	// Converting The Int Percentage Value To Float
	var percentDiff float64 = float64(diffBy) / 100.00

	differenceRange := rand.Float64()*(2*percentDiff) - percentDiff // This Generates A Random Number In The Range [-percentDiff, +percentDiff]

	newStockPrice := currStockPrice * (1 + differenceRange)
	return newStockPrice

}
