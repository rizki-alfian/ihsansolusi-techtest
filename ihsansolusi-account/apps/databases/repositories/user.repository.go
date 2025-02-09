package repositories

import (
	"gorm.io/gorm"
	"ihsansolusi-account/apps/databases/models"
	"github.com/shopspring/decimal"
	"log"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	GetUserByAccountNumber(account_number string) (*models.User, error)
	UpdateBalance(id int, amount decimal.Decimal) error
	FindUserByNIKOrPhone(nik string, phone string) (*models.User, error)
	CountUsersByAccountPrefix(prefix string) (int64, error)
	IsAccountNumberExists(accountNumber string) bool
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Hybrid: Bisa menerima seluruh struct User atau hanya sebagian field
func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByAccountNumber(account_number string) (*models.User, error) {
	var user models.User
	err := r.db.Where("account_number = ?", account_number).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}


func (r *userRepository) UpdateBalance(id int, amount decimal.Decimal) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("balance", amount).Error
}

func (r *userRepository) FindUserByNIKOrPhone(nik string, phone string) (*models.User, error) {
	var user models.User
	err := r.db.Where("nik = ? OR phone = ?", nik, phone).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) IsAccountNumberExists(accountNumber string) bool {
	var count int64
	err := r.db.Model(&models.User{}).Where("account_number = ?", accountNumber).Count(&count).Error
	if err != nil {
		log.Println("Error checking account number:", err)
		return true
	}

	return count > 0
}

func (r *userRepository) CountUsersByAccountPrefix(prefix string) (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).
		Where("account_number LIKE ?", prefix+"%").
		Count(&count).Error

	return count, err
}