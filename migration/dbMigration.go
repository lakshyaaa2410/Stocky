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
		{StockSymbol: "Reliance", StockPrice: 1000.0},
		{StockSymbol: "Apple", StockPrice: 1503.436},
		{StockSymbol: "Wipro", StockPrice: 500.46},
		{StockSymbol: "Infosys", StockPrice: 1200.63},
		{StockSymbol: "Amazon", StockPrice: 1900.35},
		{StockSymbol: "Zoho", StockPrice: 1234.567},
	}

	// Traversing The Prices "Slice" And Creating A Entry In The Database.
	for _, prices := range stockPrices {
		response := initializers.DB.Save(&prices)

		if response.Error != nil {
			logrus.Error("Error While Seeding Stock Prices")
		} else {
			logrus.Info("Stock Prices Seeded")
		}
	}
}
