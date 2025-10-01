package models

import (
	"time"

	"gorm.io/gorm"
)

type StockPrice struct {
	gorm.Model

	StockSymbol    string    `json:"stockSymbol" gorm:"not null;uniqueIndex"`
	StockPrice     float64   `json:"stockPrice" gorm:"not null;type:numeric(18,4)"`
	PriceUpdatedAt time.Time `json:"priceUpdatedAt" gorm:"autoUpdateTime"`
}
