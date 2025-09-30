package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AddReward(ginCtx *gin.Context) {

}

func StockRewardsToday(ginCtx *gin.Context) {

	userIdStr := ginCtx.Param("userId")
	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		logrus.Error("Error Pasring User ID")
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   userId,
	})

}
