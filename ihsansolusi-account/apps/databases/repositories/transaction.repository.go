package repositories

import (
	"gorm.io/gorm"
	"ihsansolusi-account/apps/databases/models"
	"github.com/shopspring/decimal"
)

type TransactionRepository interface {
	ProcessTransaction(user *models.User, amount decimal.Decimal, txType string) error
	ProcessTransfer(sender *models.User, receiver *models.User, amount decimal.Decimal) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) ProcessTransaction(user *models.User, amount decimal.Decimal, txType string) error {
	tx := r.db.Begin()

	if txType == "withdraw" {
		user.Balance = user.Balance.Sub(amount)
	} else {
		user.Balance = user.Balance.Add(amount)
	}

	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	transactionRecord := models.Transaction{
		UserID: user.ID,
		Amount: amount,
		Type:   txType,
	}
	if err := tx.Create(&transactionRecord).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *transactionRepository) ProcessTransfer(sender *models.User, receiver *models.User, amount decimal.Decimal) error {
	tx := r.db.Begin()

	sender.Balance = sender.Balance.Sub(amount)
	if err := tx.Save(sender).Error; err != nil {
		tx.Rollback()
		return err
	}

	receiver.Balance = receiver.Balance.Add(amount)
	if err := tx.Save(receiver).Error; err != nil {
		tx.Rollback()
		return err
	}

	transferRecord := models.Transaction{
		SenderID:   sender.ID,
		ReceiverID: receiver.ID,
		Amount:     amount,
		Type:       "transfer",
	}
	if err := tx.Create(&transferRecord).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
