# Stocky

func AddReward(ginCtx *gin.Context) {

	var reward models.Reward

	// 1. Binding The Request Body Into JSON
	err := ginCtx.BindJSON(&reward)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	// Utility Function To Convert Current Time Into Rounded Of Format (IST)
	reward.RewardedAt = utilities.GetDayAndTime()

	fmt.Println(reward.RewardedAt)

	// Checking If The Reward Already Exists With The Same Request Body.
	var existing models.Reward

	record := initializers.DB.Where("user_id = ? and stock_symbol = ? and action_type = ? and rewarded_at >= ?", reward.UserID, reward.StockSymbol, reward.Action, reward.RewardedAt).First(&existing)

	if record.Error == nil {
		ginCtx.JSON(http.StatusConflict, gin.H{
			"status":  "failed",
			"message": "Reward Already Exists",
		})
		return
	}

	response := initializers.DB.Create(&reward)
	if response.Error != nil {
		logrus.Error(response.Error)
		return
	}

	ginCtx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   reward,
	})
}
