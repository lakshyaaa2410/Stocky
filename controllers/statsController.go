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

type TotalSharesTodays struct {
	StockSymbol string  `json:"stockSymbol" gorm:"column:stock_symbol"`
	Shares      float64 `json:"shares" gorm:"column:shares"`
}

type TotalValuation struct {
	Valuation float64 `json:"valuation" gorm:"column:portfolio_value"`
}

func GetUserStats(ginCtx *gin.Context) {

	userIdStr := ginCtx.Param("userId")
	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		logrus.Error("Error Parsing UserId")
		return
	}

	var sharesToday []TotalSharesTodays
	var valuation []TotalValuation

	err1 := getSharesRewardedToday(userId, &sharesToday)
	err2 := getPortfolioValue(userId, &valuation)

	if err1 != nil {
		fmt.Println(err1.Error())
		return
	} else if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"status":              "success",
		"sharesRewardedToday": sharesToday,
		"portfolio":           valuation,
	})

}

func getSharesRewardedToday(userId int, sharesRewardedToday *[]TotalSharesTodays) error {

	_, today := utilities.GetDayAndTime()

	query := `SELECT rewardTable.stock_symbol AS stock_symbol,
    		SUM(rewardTable.shares) AS Shares,
    		ROUND(SUM(rewardTable.shares * stockPriceTable.stock_price), 2) AS Value
			FROM rewards rewardTable
			JOIN stock_prices stockPriceTable 
    		ON rewardTable.stock_symbol = stockPriceTable.stock_symbol
			WHERE rewardTable.user_id = ?
  			AND DATE(rewardTable.rewarded_at) = ?
			GROUP BY rewardTable.stock_symbol`

	return initializers.DB.Raw(query, userId, today).Scan(&sharesRewardedToday).Error
}

func getPortfolioValue(userId int, portfolioValue *[]TotalValuation) error {

	query := `SELECT ROUND(SUM(rewardsTable.shares * stockPricesTable.stock_price),2) AS portfolio_value
			FROM rewards rewardsTable
			JOIN stock_prices stockPricesTable ON rewardsTable.stock_symbol = stockPricesTable.stock_symbol
			WHERE rewardsTable.user_id = ?`

	return initializers.DB.Raw(query, userId).Scan(&portfolioValue).Error
}
