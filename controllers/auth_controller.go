package controllers

import (
	"github.com/gin-gonic/gin"
	"go-product/domain/dto"
	"go-product/domain/entity"
	"go-product/helpers"
	"go-product/repositories/i_repositories"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type authController struct {
	userRepository i_repositories.UserRepository
}

type AuthController interface {
	HandleUserLogin(c *gin.Context)
	HandleUserRegister(c *gin.Context)
}

func NewAuthController(userRepository i_repositories.UserRepository) AuthController {
	return &authController{userRepository}
}

func (ctr authController) HandleUserLogin(c *gin.Context) {
	payload := c.MustGet("payload").(dto.AuthenticationRequest)

	user, err := ctr.userRepository.GetUserByEmail(payload.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ApiResponse{
				Code:    http.StatusUnauthorized,
				Status:  "UNAUTHORIZED",
				Message: "Invalid email or password. Please check your email and password and try again",
			})
			return
		}

		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ApiResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "An error occurred while processing your request. Please try again later",
		})
		return
	}

	isValidPassword := helpers.ValidatePassword(user.Password, payload.Password)
	if !isValidPassword {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ApiResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "Invalid email or password. Please check your email and password and try again",
		})
		return
	}

	token, err := helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ApiResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "An error occurred while processing your request. Please try again later",
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Login successful. Please use the token to access protected resources",
		Data: dto.AuthenticationResponse{
			Token: token,
		},
	})
}

func (ctr authController) HandleUserRegister(c *gin.Context) {
	payload := c.MustGet("payload").(dto.UserRequest)

	var user entity.User
	user.FromRequest(payload)

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ApiResponse{
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
			c.AbortWithStatusJSON(http.StatusConflict, dto.ApiResponse{
				Code:    http.StatusConflict,
				Status:  "CONFLICT",
				Message: "User with the same email already exists. Please use a different email and try again",
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ApiResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "User created successfully",
		Data:    result.ToResponse(),
	})
}
