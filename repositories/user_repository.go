package repositories

import (
	"go-product/domain/entity"
	"go-product/pkg/errors"
	"go-product/repositories/repo_interfaces"
	"gorm.io/gorm"
	"log"
)

type userRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) repo_interfaces.UserRepository {
	return &userRepository{db}
}

func (u userRepository) CreateUser(user *entity.User) (*entity.User, errors.Error) {
	err := u.DB.Model(&entity.User{}).Create(user).Error
	if err != nil {
		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return user, nil
}

func (u userRepository) GetUserByEmail(email string) (*entity.User, errors.Error) {
	var user entity.User

	err := u.DB.Model(&entity.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errs := errors.NewUnauthenticatedError("Invalid email or password. Please check your email and password and try again")
			return nil, errs
		}

		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return &user, nil
}
