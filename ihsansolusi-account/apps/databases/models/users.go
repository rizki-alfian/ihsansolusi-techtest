package models

import (
	"gorm.io/gorm"
	"github.com/shopspring/decimal"
)

type User struct {
	gorm.Model
	Name      		string  		`json:"name,omitempty"`
	NIK      		string  		`json:"nik" gorm:"unique;not null"`
	Phone    		string  		`json:"phone" gorm:"unique;not null"`
	AccountNumber 	string 			`json:"account_number" gorm:"unique;not null"`
	Balance       	decimal.Decimal `json:"balance" gorm:"type:decimal(10,2);default:0"`
}