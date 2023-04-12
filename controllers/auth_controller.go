package controllers

import (
	"github.com/gin-gonic/gin"
	"go-product/domain/dto"
	"go-product/services"
	"net/http"
)

type authController struct {
	authService services.AuthService
}

type AuthController interface {
	HandleUserLogin(c *gin.Context)
	HandleUserRegister(c *gin.Context)
}

func NewAuthController(authService services.AuthService) AuthController {
	return &authController{authService}
}

func (ctr authController) HandleUserLogin(c *gin.Context) {
	payload := c.MustGet("payload").(dto.AuthenticationRequest)
	response, err := ctr.authService.Login(payload)
	if err != nil {
		c.AbortWithStatusJSON(err.Code(), dto.ApiResponse{
			Code:    err.Code(),
			Status:  err.Status(),
			Message: err.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Login successful. Please use the token to access protected resources",
		Data:    response,
	})
}

func (ctr authController) HandleUserRegister(c *gin.Context) {
	payload := c.MustGet("payload").(dto.UserRequest)
	response, err := ctr.authService.Register(payload)
	if err != nil {
		c.AbortWithStatusJSON(err.Code(), dto.ApiResponse{
			Code:    err.Code(),
			Status:  err.Status(),
			Message: err.Message(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.ApiResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "User created successfully",
		Data:    response,
	})
}
