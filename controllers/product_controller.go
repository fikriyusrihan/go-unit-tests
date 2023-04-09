package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-product/domain/dto"
	"go-product/services"
	"net/http"
	"strconv"
)

type productController struct {
	productService services.ProductService
}

type ProductController interface {
	HandleCreateProduct(c *gin.Context)
	HandleUpdateProduct(c *gin.Context)
	HandleDeleteProduct(c *gin.Context)
	HandleGetAllProduct(c *gin.Context)
	HandleGetProductById(c *gin.Context)
}

func NewProductController(
	productService services.ProductService,
) ProductController {
	return &productController{productService}
}

func (p productController) HandleCreateProduct(c *gin.Context) {
	payload := c.MustGet("payload").(dto.ProductRequest)
	claim := c.MustGet("claim").(jwt.MapClaims)
	uid := uint(claim["id"].(float64))

	response, err := p.productService.CreateProduct(uid, payload)
	if err != nil {
		c.AbortWithStatusJSON(err.Code(), dto.ApiResponse{
			Code:    err.Code(),
			Status:  err.Status(),
			Message: err.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "Product created successfully",
		Data:    response,
	})
}

func (p productController) HandleUpdateProduct(c *gin.Context) {
	payload := c.MustGet("payload").(dto.ProductRequest)

	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Invalid product id",
		})
		return
	}

	response, errs := p.productService.UpdateProduct(productId, payload)
	if errs != nil {
		c.AbortWithStatusJSON(errs.Code(), dto.ApiResponse{
			Code:    errs.Code(),
			Status:  errs.Status(),
			Message: errs.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Product updated successfully",
		Data:    response,
	})
}

func (p productController) HandleDeleteProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Invalid product id",
		})
		return
	}

	errs := p.productService.DeleteProduct(productId)
	if errs != nil {
		c.AbortWithStatusJSON(errs.Code(), dto.ApiResponse{
			Code:    errs.Code(),
			Status:  errs.Status(),
			Message: errs.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Product deleted successfully",
	})
}

func (p productController) HandleGetAllProduct(c *gin.Context) {
	response, err := p.productService.GetProducts()
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
		Message: "Product list retrieved successfully",
		Data:    response,
	})
}

func (p productController) HandleGetProductById(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Invalid product id",
		})
		return
	}

	response, errs := p.productService.GetProductById(productId)
	if errs != nil {
		c.AbortWithStatusJSON(errs.Code(), dto.ApiResponse{
			Code:    errs.Code(),
			Status:  errs.Status(),
			Message: errs.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Product retrieved successfully",
		Data:    response,
	})
}
