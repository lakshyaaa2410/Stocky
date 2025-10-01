package models

import (
	"time"

	"gorm.io/gorm"
)

type Reward struct {
	gorm.Model
	UserID      uint      `json:"userId" gorm:"not null"`
	StockSymbol string    `json:"stockSymbol" gorm:"size:20;not null"`
	Action      string    `json:"action" gorm:"not null"`
	Shares      float64   `json:"shares" gorm:"not null"`
	RewardedAt  time.Time `json:"rewardedAt" gorm:"not null"`
}
