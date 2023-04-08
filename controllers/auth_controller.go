package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-product/domain"
	"go-product/helpers"
	"go-product/repository"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type AuthController interface {
	HandleUserLogin(c *gin.Context)
	HandleUserRegister(c *gin.Context)
}

type authController struct {
	userRepository repository.UserRepository
}

func NewAuthController(userRepository repository.UserRepository) AuthController {
	return &authController{userRepository}
}

func (ctr authController) HandleUserLogin(c *gin.Context) {
	payload := c.MustGet("payload").(domain.AuthenticationRequest)

	user, err := ctr.userRepository.GetUserByEmail(payload.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ApiResponse{
				Code:    http.StatusUnauthorized,
				Status:  "UNAUTHORIZED",
				Message: "Invalid email or password. Please check your email and password and try again",
			})
		}

		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.ApiResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "An error occurred while processing your request. Please try again later",
		})
	}

	fmt.Println(user.Password, payload.Password)
	isValidPassword := helpers.ValidatePassword(user.Password, payload.Password)
	if !isValidPassword {
		c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ApiResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "Invalid email or password. Please check your email and password and try again",
		})
		return
	}

	token, err := helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.ApiResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "An error occurred while processing your request. Please try again later",
		})
	}

	c.JSON(http.StatusOK, domain.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Login successful. Please use the token to access protected resources",
		Data: domain.AuthenticationResponse{
			Token: token,
		},
	})
}

func (ctr authController) HandleUserRegister(c *gin.Context) {
	payload := c.MustGet("payload").(domain.UserRequest)

	var user domain.User
	user.FromRequest(payload)

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.ApiResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "An error occurred while processing your request. Please try again later",
		})
		return
	}

	user.Password = hashedPassword
	result, err := ctr.userRepository.CreateUser(&user)
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			c.AbortWithStatusJSON(http.StatusConflict, domain.ApiResponse{
				Code:    http.StatusConflict,
				Status:  "CONFLICT",
				Message: "User with the same email already exists. Please use a different email and try again",
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.ApiResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, domain.ApiResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "User created successfully",
		Data:    result.ToResponse(),
	})
}
