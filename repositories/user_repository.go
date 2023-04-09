package repositories

import (
	"go-product/domain/entity"
	"go-product/repositories/i_repositories"
	"gorm.io/gorm"
)

type userRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) i_repositories.UserRepository {
	return &userRepository{db}
}

func (u userRepository) CreateUser(user *entity.User) (*entity.User, error) {
	err := u.DB.Model(&entity.User{}).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := u.DB.Model(&entity.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
