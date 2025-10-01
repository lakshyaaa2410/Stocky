package models

import (
	"time"

	"gorm.io/gorm"
)

type Reward struct {
	gorm.Model
	UserID      uint      `json:"userId" gorm:"not null" binding:"required"`
	StockSymbol string    `json:"stockSymbol" gorm:"size:20;not null" binding:"required"`
	Action      string    `json:"action" gorm:"not null" binding:"required"`
	Shares      float64   `json:"shares" gorm:"not null" binding:"required"`
	RewardedAt  time.Time `json:"rewardedAt" gorm:"not null"`
}
