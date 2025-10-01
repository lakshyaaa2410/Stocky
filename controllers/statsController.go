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

	err = getSharesRewardedToday(userId, &sharesToday)
	err1 := getPortfolioValue(userId, &valuation)

	if err != nil {
		fmt.Println(err.Error())
		return
	} else if err1 != nil {
		fmt.Println(err1.Error())
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"sharesToday": sharesToday,
		"portfolio":   valuation,
	})

}

func getSharesRewardedToday(userId int, sharesRewardedToday *[]TotalSharesTodays) error {

	_, today := utilities.GetDayAndTime()

	query := `SELECT stock_symbol, CAST(SUM(shares) AS DOUBLE PRECISION) AS shares 
			FROM rewards 
			WHERE user_id = ? AND DATE(rewarded_at) = ? 
			GROUP BY stock_symbol`

	return initializers.DB.Raw(query, userId, today).Scan(&sharesRewardedToday).Error
}

func getPortfolioValue(userId int, portfolioValue *[]TotalValuation) error {
	query := `SELECT 
    		COALESCE(SUM(rewardsTable.shares * stockPriceTable.stock_price), 0) AS portfolio_value
			FROM rewards rewardsTable
			JOIN stock_prices stockPriceTable ON rewardsTable.stock_symbol = stockPriceTable.stock_symbol
			WHERE rewardsTable.user_id = ?`

	return initializers.DB.Raw(query, userId).Scan(&portfolioValue).Error
}
