package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-product/domain"
	"go-product/repository"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type productController struct {
	userRepository    repository.UserRepository
	productRepository repository.ProductRepository
}

type ProductController interface {
	HandleCreateProduct(c *gin.Context)
	HandleUpdateProduct(c *gin.Context)
	HandleDeleteProduct(c *gin.Context)
	HandleGetAllProduct(c *gin.Context)
	HandleGetProductById(c *gin.Context)
}

func NewProductController(
	userRepository repository.UserRepository,
	productRepository repository.ProductRepository,
) ProductController {
	return &productController{userRepository, productRepository}
}

func (p productController) HandleCreateProduct(c *gin.Context) {
	payload := c.MustGet("payload").(domain.ProductRequest)
	claim := c.MustGet("claim").(jwt.MapClaims)

	var product domain.Product
	product.FromRequest(payload)

	uid := uint(claim["id"].(float64))
	product.UserID = uid

	result, err := p.productRepository.CreateProduct(&product)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.ApiResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "An error occurred while processing your request. Please try again later",
		})
		return
	}

	c.JSON(http.StatusOK, domain.ApiResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "Product created successfully",
		Data:    result.ToResponse(),
	})
}

func (p productController) HandleUpdateProduct(c *gin.Context) {
	payload := c.MustGet("payload").(domain.ProductRequest)

	var product domain.Product
	product.FromRequest(payload)

	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Invalid product id",
		})
		return
	}

	result, err := p.productRepository.UpdateProduct(productId, &product)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, domain.ApiResponse{
				Code:    http.StatusNotFound,
				Status:  "NOT_FOUND",
				Message: "Product not found",
			})
			return
		}

		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.ApiResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "An error occurred while processing your request. Please try again later",
		})
		return
	}

	c.JSON(http.StatusOK, domain.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Product updated successfully",
		Data:    result.ToResponse(),
	})
}

func (p productController) HandleDeleteProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Invalid product id",
		})
		return
	}

	err = p.productRepository.DeleteProduct(productId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, domain.ApiResponse{
				Code:    http.StatusNotFound,
				Status:  "NOT_FOUND",
				Message: "Product not found",
			})
			return
		}

		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, domain.ApiResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "An error occurred while processing your request. Please try again later",
		})
		return
	}

	c.JSON(http.StatusOK, domain.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Product deleted successfully",
	})
}

func (p productController) HandleGetAllProduct(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p productController) HandleGetProductById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
