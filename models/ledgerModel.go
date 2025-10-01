package models

import (
	"time"

	"gorm.io/gorm"
)

type Ledger struct {
	gorm.Model

	UserID          uint      `json:"userId" gorm:"index;not null"`
	RewardID        uint      `json:"rewardId" gorm:"not null"`
	TransactionType string    `json:"transactionType" gorm:"not null"`
	Amount          float64   `json:"amount" gorm:"not null;type:numeric(18,4)"`
	AmountUnit      string    `json:"amountUnit" gorm:"not null"`
	FlowType        string    `json:"flowType" gorm:"not null"`
	Action          string    `json:"action" gorm:"not null"`
	Account         string    `json:"account" gorm:"not null"`
	EnteredAt       time.Time `json:"enteredAt" gorm:"autoUpdateTime"`
}
