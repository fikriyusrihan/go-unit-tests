package services

import (
	"go-product/domain/dto"
	"go-product/domain/entity"
	"go-product/pkg/errors"
	"go-product/pkg/helpers"
	"go-product/repositories/i_repositories"
	"gorm.io/gorm"
	"log"
)

type AuthService interface {
	Login(payload dto.AuthenticationRequest) (*dto.AuthenticationResponse, errors.Error)
	Register(payload dto.UserRequest) (*dto.UserResponse, errors.Error)
}

type authService struct {
	userRepository i_repositories.UserRepository
}

func NewAuthService(userRepository i_repositories.UserRepository) AuthService {
	return &authService{userRepository}
}

func (a authService) Login(request dto.AuthenticationRequest) (*dto.AuthenticationResponse, errors.Error) {
	user, err := a.userRepository.GetUserByEmail(request.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			httpError := errors.NewUnauthenticatedError("Invalid email or password. Please check your email and password and try again")
			return nil, httpError
		}

		log.Println(err)
		httpError := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, httpError
	}

	isValidPassword := helpers.ValidatePassword(user.Password, request.Password)
	if !isValidPassword {
		httpError := errors.NewUnauthenticatedError("Invalid email or password. Please check your email and password and try again")
		return nil, httpError
	}

	token, err := helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Println(err)
		httpError := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, httpError
	}

	response := &dto.AuthenticationResponse{
		Token: token,
	}

	return response, nil
}

func (a authService) Register(payload dto.UserRequest) (*dto.UserResponse, errors.Error) {
	var user entity.User
	user.FromRequest(payload)

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		httpError := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, httpError
	}

	user.Password = hashedPassword
	result, err := a.userRepository.CreateUser(&user)
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			httpError := errors.NewConflictError("Email already exists")
			return nil, httpError
		}

		log.Println(err)
		httpError := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, httpError
	}

	response := result.ToResponse()
	return &response, nil
}
