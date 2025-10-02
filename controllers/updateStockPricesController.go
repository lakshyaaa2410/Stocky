package controllers

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lakshyaaa2410/stocky/initializers"
	"github.com/lakshyaaa2410/stocky/models"
	"github.com/sirupsen/logrus"
)

func UpdateStockPrices(ginCtx *gin.Context) {

	// Variable To Store Rewards
	var stockPrices []models.StockPrice

	// Fetching Current Stock Prices From Database.
	err := initializers.DB.Find(&stockPrices).Error

	// Checking For Potential Errors
	if err != nil {
		logrus.Error("Error Fetching Stock Prices")
		return
	}

	// Looping Thorugh The Result
	for _, stockPrices := range stockPrices {
		// Getting The Current Stock Price
		currentStockPrice := stockPrices.StockPrice

		// Helper Method To Get Randomly Generated New Price
		// Arguments - (current Price, % Change In The New Price)
		newStockPrice := getRandomStockPrice(currentStockPrice, 10)

		// Updating The Prices In The Database
		err := initializers.DB.Model(&stockPrices).Update("StockPrice", newStockPrice).Error

		// Checking For Potential Errors
		if err != nil {
			logrus.Error("Error Updating Price")
			continue
		}
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Stock Prices Updated",
	})

}

func getRandomStockPrice(currStockPrice float64, diffBy int) float64 {

	// Converting The Int Percentage Value To Float
	var percentDiff float64 = float64(diffBy) / 100.00

	// This Generates A Random Number In The Range [-percentDiff, +percentDiff]
	differenceRange := rand.Float64()*(2*percentDiff) - percentDiff

	newStockPrice := currStockPrice * (1 + differenceRange)
	return newStockPrice

}
