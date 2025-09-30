package models

import (
	"time"

	"gorm.io/gorm"
)

type Reward struct {
	gorm.Model

	UserId      int       `json:"userId"`
	StockSymbol string    `json:"stockSymbol"`
	Shares      float32   `json:"shares"`
	RewardedAt  time.Time `json:"rewardedAt"`
}
