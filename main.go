package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lakshyaaa2410/stocky/controllers"
	"github.com/lakshyaaa2410/stocky/initializers"
	"github.com/sirupsen/logrus"
)

func init() {

	// Helper Function To Load Enviornment Variables
	initializers.LoadEnvVariables()

	// Connect To Database
	initializers.ConnectDB()
}

func main() {

	// Initializing Gin Router
	router := gin.Default()

	// POST Method To Add Rewards For A User
	router.POST("/reward", controllers.AddReward)

	// GET Method To Fetch All The Records Of Today's Stock.
	router.GET("/today-stocks/:userId", controllers.GetStockRewardsToday)

	// GET Method To Fetch User Stats
	router.GET("/stats/:userId", controllers.GetUserStats)

	// GET Method To Fetch Stock History Of An User
	router.GET("/historical-inr/:userId", controllers.GetStockHistory)

	// PUT Method For Manual Trigger For Stocks Price Update
	router.PUT("/update-stock-prices", controllers.UpdateStockPrices)

	// GET Method To Fetch User's Portfolio
	router.GET("/portfolio/:userId", controllers.GetUserPortfolio)

	// The Server Runs On The Port Number Specified In .env File.
	// If No Value Is Present, It Runs On 8080 By Default (8080 Should Not Be Occupied)
	err := router.Run()
	if err != nil {
		logrus.Error("Error Starting The Server")
	}

}
