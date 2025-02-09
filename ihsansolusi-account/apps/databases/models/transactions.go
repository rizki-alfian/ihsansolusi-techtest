package models

import (
	"gorm.io/gorm"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	gorm.Model
	UserID    	uint    `json:"user_id"` //For deposit/withdraw
	SenderID   	uint   	`json:"sender_id,omitempty"` // For transfer
	ReceiverID 	uint   	`json:"receiver_id,omitempty"` // For transfer
	Amount      decimal.Decimal `json:"amount" gorm:"type:decimal(10,2); not null"`
	Type      	string  `json:"type"` // "deposit", "withdraw", "transfer"
}