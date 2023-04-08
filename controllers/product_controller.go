package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-product/domain"
	"go-product/repository"
	"log"
	"net/http"
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
	}

	c.JSON(http.StatusOK, domain.ApiResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "Product created successfully",
		Data:    result.ToResponse(),
	})
}

func (p productController) HandleUpdateProduct(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p productController) HandleDeleteProduct(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p productController) HandleGetAllProduct(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p productController) HandleGetProductById(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
