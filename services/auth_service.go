package services

import (
	"go-product/domain/dto"
	"go-product/domain/entity"
	"go-product/pkg/errors"
	"go-product/pkg/helpers"
	"go-product/repositories/repo_interfaces"
	"log"
)

type AuthService interface {
	Login(payload dto.AuthenticationRequest) (*dto.AuthenticationResponse, errors.Error)
	Register(payload dto.UserRequest) (*dto.UserResponse, errors.Error)
}

type authService struct {
	userRepository repo_interfaces.UserRepository
}

func NewAuthService(userRepository repo_interfaces.UserRepository) AuthService {
	return &authService{userRepository}
}

func (a authService) Login(request dto.AuthenticationRequest) (*dto.AuthenticationResponse, errors.Error) {
	user, errs := a.userRepository.GetUserByEmail(request.Email)
	if errs != nil {
		return nil, errs
	}

	isValidPassword := helpers.ValidatePassword(user.Password, request.Password)
	if !isValidPassword {
		newErrs := errors.NewUnauthenticatedError("Invalid email or password. Please check your email and password and try again")
		return nil, newErrs
	}

	token, err := helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Println(err)
		newErrs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, newErrs
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
		newErrs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, newErrs
	}

	user.Password = hashedPassword
	result, errs := a.userRepository.CreateUser(&user)
	if errs != nil {
		return nil, errs
	}

	response := result.ToResponse()
	return &response, nil
}
