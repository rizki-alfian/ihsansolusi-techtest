package transactions

import (
	"errors"
	"ihsansolusi-account/apps/databases/repositories"
	"github.com/shopspring/decimal"
)

type TransactionService interface {
	ProcessTransaction(accountNumber string, amount float64, txType string) error
	ProcessTransfer(senderID string, receiverID string, amount float64) error
}

type transactionService struct {
	userRepo       repositories.UserRepository
	transactionRepo repositories.TransactionRepository
}

func NewTransactionService(userRepo repositories.UserRepository, transactionRepo repositories.TransactionRepository) *transactionService {
	return &transactionService{userRepo: userRepo, transactionRepo: transactionRepo}
}

func (s *transactionService) ProcessTransaction(accountNumber string, amount float64, txType string) error {
	user, err := s.userRepo.GetUserByAccountNumber(accountNumber)
	if err != nil {
		return errors.New("User not found")
	}

	if user == nil {
		return errors.New("User not found")
	}

	if txType == "withdraw" && user.Balance.LessThan(decimal.NewFromFloat(amount)) {
		return errors.New("Insufficient balance")
	}

	decimalAmount := decimal.NewFromFloat(amount)

	err = s.transactionRepo.ProcessTransaction(user, decimalAmount, txType)
	if err != nil {
		return errors.New("Transaction failed: " + err.Error())
	}

	return nil
}

func (s *transactionService) ProcessTransfer(senderID string, receiverID string, amount float64) error {
	sender, err := s.userRepo.GetUserByAccountNumber(senderID)
	if err != nil {
		return errors.New("Sender not found")
	}

	if sender == nil {
		return errors.New("Sender not found")
	}

	receiver, err := s.userRepo.GetUserByAccountNumber(receiverID)
	if err != nil {
		return errors.New("Receiver not found")
	}

	if receiver == nil {
		return errors.New("Receiver not found")
	}

	decimalAmount := decimal.NewFromFloat(amount)

	if sender.Balance.LessThan(decimalAmount) {
		return errors.New("Insufficient balance")
	}

	err = s.transactionRepo.ProcessTransfer(sender, receiver, decimalAmount)
	if err != nil {
		return errors.New("Transfer failed: " + err.Error())
	}

	return nil
}
