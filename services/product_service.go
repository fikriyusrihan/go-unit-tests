package services

import (
	"go-product/domain/dto"
	"go-product/domain/entity"
	"go-product/pkg/errors"
	"go-product/repositories/repo_interfaces"
	"log"
)

type ProductService interface {
	CreateProduct(uid uint, payload dto.ProductRequest) (*dto.ProductResponse, errors.Error)
	UpdateProduct(pid int, payload dto.ProductRequest) (*dto.ProductResponse, errors.Error)
	DeleteProduct(pid int) errors.Error
	GetProducts() ([]dto.ProductResponse, errors.Error)
	GetProductById(pid int) (*dto.ProductResponse, errors.Error)
}

type productService struct {
	productRepository repo_interfaces.ProductRepository
}

func NewProductService(productRepository repo_interfaces.ProductRepository) ProductService {
	return &productService{productRepository}
}

func (p productService) CreateProduct(uid uint, payload dto.ProductRequest) (*dto.ProductResponse, errors.Error) {
	var product entity.Product
	product.FromRequest(payload)
	product.UserID = uid

	result, err := p.productRepository.CreateProduct(&product)
	if err != nil {
		log.Println(err)
		httpError := errors.NewInternalServerError("An error occurred while creating product. Please try again later.")
		return nil, httpError
	}

	response := result.ToResponse()
	return &response, nil
}

func (p productService) UpdateProduct(pid int, payload dto.ProductRequest) (*dto.ProductResponse, errors.Error) {
	var product entity.Product
	product.FromRequest(payload)

	result, err := p.productRepository.UpdateProduct(pid, &product)
	if err != nil {
		return nil, err
	}

	response := result.ToResponse()
	return &response, nil
}

func (p productService) DeleteProduct(pid int) errors.Error {
	err := p.productRepository.DeleteProduct(pid)
	if err != nil {
		return err
	}

	return nil
}

func (p productService) GetProducts() ([]dto.ProductResponse, errors.Error) {
	result, err := p.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}

	var response []dto.ProductResponse
	for _, product := range result {
		response = append(response, product.ToResponse())
	}

	return response, nil
}

func (p productService) GetProductById(pid int) (*dto.ProductResponse, errors.Error) {
	result, err := p.productRepository.GetProductById(pid)
	if err != nil {
		return nil, err
	}

	response := result.ToResponse()
	return &response, nil
}
