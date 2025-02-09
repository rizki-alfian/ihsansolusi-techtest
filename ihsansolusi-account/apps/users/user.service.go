package users

import (
	"ihsansolusi-account/apps/databases/models"
	"ihsansolusi-account/apps/databases/repositories"
	"ihsansolusi-account/apps/cores/utils"
	"github.com/shopspring/decimal"
	"errors"
)

type UserService interface {
	RegisterUser(name string, nik string, phone string) (*models.User, error)
	CheckBalance(account_number string) (decimal.Decimal, error)
	GetUserByNIKOrPhone(nik string, phone string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(name string, nik string, phone string) (*models.User, error) {
	half_account_number := utils.HalfAccountNumber("62")
	last_prefix, err := s.repo.CountUsersByAccountPrefix(half_account_number)

	if err != nil {
		return nil, err
	}

	if last_prefix == 0 {
		last_prefix = 0
	} 
	last_prefix += 1

	account_number := utils.GenerateAccountNumber("62", int(last_prefix))

	user := &models.User{Name: name, NIK: nik, Phone: phone, AccountNumber: account_number}
	return s.repo.CreateUser(user)
}

func (s *userService) CheckBalance(account_number string) (decimal.Decimal, error) {
	user, err := s.repo.GetUserByAccountNumber(account_number)
	if err != nil {
		return decimal.Zero, err
	}

	if user == nil {
		return decimal.Zero, errors.New("User not found")
	}
	return user.Balance, nil
}

func (s *userService) GetUserByNIKOrPhone(nik string, phone string) (*models.User, error) {
	return s.repo.FindUserByNIKOrPhone(nik, phone)
}