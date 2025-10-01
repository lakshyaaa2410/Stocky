package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lakshyaaa2410/stocky/initializers"
	"github.com/lakshyaaa2410/stocky/models"
	"github.com/lakshyaaa2410/stocky/utilities"
	"github.com/sirupsen/logrus"
)

func AddReward(ginCtx *gin.Context) {

	// 1. Initializing A New Reward Variable
	var reward models.Reward

	// 2. Binding The Incoming JSON Body Into Reward Variable
	err := ginCtx.ShouldBindJSON(&reward)

	// 2.A. Checking If There Is Any Error In The Request Body.
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	var existingReward models.Reward

	// Making Sure The "Action" Is Always First Letter Capitalized
	reward.Action = utilities.UpperCaseFirstLetter(reward.Action)

	// 3. Checking If The User Has Already Onboarded Or Not
	if action := reward.Action; action == "Onboarding" {
		fmt.Println("In Onboarding Condition")

		err := initializers.DB.Where("user_id = ? AND action = ?", reward.UserID, "Onboarding").First(&existingReward).Error

		// Means A Result With Same User And "Onboarding" Status Was Found
		if err == nil {
			ginCtx.JSON(http.StatusConflict, gin.H{
				"status":  "failed",
				"message": "Onboarding Reward Already Given",
			})
			return
		}
	}

	reward.RewardedAt = utilities.GetDayAndTime()

	response := initializers.DB.Create(&reward)
	if response.Error != nil {
		ginCtx.JSON(http.StatusConflict, gin.H{
			"status":  "failed",
			"message": "Onboarding Reward Already Given",
		})

		return
	}

	ginCtx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": gin.H{
			"reward": reward,
		},
	})
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
