package repository

import (
	"go-product/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	*gorm.DB
}

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u userRepository) CreateUser(user *domain.User) (*domain.User, error) {
	err := u.DB.Model(&domain.User{}).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := u.DB.Model(&domain.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
