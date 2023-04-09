package i_repositories

import (
	"go-product/domain/entity"
	"go-product/pkg/errors"
)

type UserRepository interface {
	CreateUser(user *entity.User) (*entity.User, errors.Error)
	GetUserByEmail(email string) (*entity.User, errors.Error)
}
