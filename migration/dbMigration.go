package main

import (
	"fmt"

	"github.com/lakshyaaa2410/stocky/initializers"
	"github.com/lakshyaaa2410/stocky/models"
	"github.com/sirupsen/logrus"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {

	err := initializers.DB.AutoMigrate(&models.Reward{}, &models.StockPrice{}, &models.Ledger{})

	if err != nil {
		logrus.Error("Error Migrating Models")
		return
	}
	fmt.Println("Migration Done")

	// Seeding The Stock Prices Initially
	seedingStockPrices()
}

func seedingStockPrices() {

	// Cretaing Dummy Stock Prices For Seeding.
	stockPrices := []models.StockPrice{
		{StockSymbol: "RELIANCE", StockPrice: 2500.00},
		{StockSymbol: "TCS", StockPrice: 1500.00},
		{StockSymbol: "WIPRO", StockPrice: 500.00},
		{StockSymbol: "INFOSYS", StockPrice: 1200.00},
	}

	// Traversing The Prices Slice And Creating A Entry In The Database.
	for _, prices := range stockPrices {
		response := initializers.DB.Create(&prices)

		if response.Error != nil {
			logrus.Error("Error While Seeding Stock Prices")
		} else {
			logrus.Info("Stock Prices Seeded")
		}
	}
}
