package controllers

import (
	"fmt"
	"net/http"

	// "net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lakshyaaa2410/stocky/initializers"
	"github.com/lakshyaaa2410/stocky/utilities"
	"github.com/sirupsen/logrus"
)

// Struct To Store The Stock Symbol And Shares For Today.
type TotalSharesTodays struct {
	StockSymbol string  `json:"stockSymbol" gorm:"column:stock_symbol"`
	Shares      float64 `json:"shares" gorm:"column:shares"`
}

// Struct To Store The Total Amount.
type TotalValuation struct {
	Valuation float64 `json:"valuation" gorm:"column:portfolio_value"`
}

func GetUserStats(ginCtx *gin.Context) {
	// Converting UserID From String To Int
	userIdStr := ginCtx.Param("userId")
	userId, err := strconv.Atoi(userIdStr)

	// Checking For Potential Errors During Converstioin.
	if err != nil {
		logrus.Error("Error Parsing UserId")
		return
	}

	// Variables To Store Today's Shares And Total Amount
	var sharesToday []TotalSharesTodays
	var valuation []TotalValuation

	// Helper Methods To Get Today's Rewards And Total Amount
	err1 := getSharesRewardedToday(userId, &sharesToday)
	err2 := getPortfolioValue(userId, &valuation)

	// Checking For Potential Errors After Data Fetching.
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	} else if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	// Returning The Response, If No Errors
	ginCtx.JSON(http.StatusOK, gin.H{
		"status":              "success",
		"sharesRewardedToday": sharesToday,
		"portfolio":           valuation,
	})

}

func getSharesRewardedToday(userId int, sharesRewardedToday *[]TotalSharesTodays) error {

	// Utility Helper Method To Get Today's Date.
	_, today := utilities.GetDayAndTime()

	// Creating A Raw SQL Query To Fetch User's Reward Data Of Current Day
	query := `SELECT rewardTable.stock_symbol AS stock_symbol,
    		SUM(rewardTable.shares) AS Shares,
    		ROUND(SUM(rewardTable.shares * stockPriceTable.stock_price), 2) AS Value
			FROM rewards rewardTable
			JOIN stock_prices stockPriceTable 
    		ON rewardTable.stock_symbol = stockPriceTable.stock_symbol
			WHERE rewardTable.user_id = ?
  			AND DATE(rewardTable.rewarded_at) = ?
			GROUP BY rewardTable.stock_symbol`

	// Executing The Raw Query And Saving The Result Into sharesRewardedToday Variable
	// Returning The Error (If Any)
	return initializers.DB.Raw(query, userId, today).Scan(&sharesRewardedToday).Error
}

func getPortfolioValue(userId int, portfolioValue *[]TotalValuation) error {

	// Creating A Raw SQL Query To Fetch User's Total Amount
	query := `SELECT ROUND(SUM(rewardsTable.shares * stockPricesTable.stock_price),2) AS portfolio_value
			FROM rewards rewardsTable
			JOIN stock_prices stockPricesTable ON rewardsTable.stock_symbol = stockPricesTable.stock_symbol
			WHERE rewardsTable.user_id = ?`

	// Executing The Raw Query And Saving The Result Into portfolioValue Variable
	// Returning The Error (If Any)
	return initializers.DB.Raw(query, userId).Scan(&portfolioValue).Error
}
