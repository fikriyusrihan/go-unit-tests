package i_repositories

import "go-product/domain/entity"

type UserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}
