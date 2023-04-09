package services_test

import (
	"github.com/stretchr/testify/assert"
	"go-product/domain/entity"
	"go-product/pkg/errors"
	"go-product/repositories/repositories_mock"
	"go-product/services"
	"testing"
)

var productRepository = &repositories_mock.ProductRepositoryMock{}
var productService = services.NewProductService(productRepository)

func TestProductService_GetAllProduct_Found(t *testing.T) {
	expectedProducts := []*entity.Product{
		{
			GormModel: entity.GormModel{
				ID: 1,
			},
			Title:       "Product 1",
			Description: "Product 1 Description",
		},
		{
			GormModel: entity.GormModel{
				ID: 2,
			},
			Title:       "Product 2",
			Description: "Product 2 Description",
		},
	}
	productRepository.Mock.On("GetProducts").Return(expectedProducts, nil)

	actualProducts, actualErr := productService.GetProducts()

	assert.Nil(t, actualErr, "error should be nil")
	assert.NotNil(t, actualProducts, "products should not be nil")
	assert.Equal(t, len(expectedProducts), len(actualProducts), "products length should be equal")

	assert.Equal(t, expectedProducts[0].GormModel.ID, actualProducts[0].ID, "product id should be equal")
	assert.Equal(t, expectedProducts[0].Title, actualProducts[0].Title, "product title should be equal")
	assert.Equal(t, expectedProducts[0].Description, actualProducts[0].Description, "product description should be equal")

	assert.Equal(t, expectedProducts[1].GormModel.ID, actualProducts[1].ID, "product id should be equal")
	assert.Equal(t, expectedProducts[1].Title, actualProducts[1].Title, "product title should be equal")
	assert.Equal(t, expectedProducts[1].Description, actualProducts[1].Description, "product description should be equal")
}

func TestProductService_GetAllProduct_NotFound(t *testing.T) {
	// Clear the expected calls so that the mock will not return any value
	// This is needed because the mock will return the last value that was set
	// and if the previous test set a value, this test will return that value
	productRepository.Mock.ExpectedCalls = nil

	expectedErr := errors.NewNotFoundError("Products not found")
	productRepository.Mock.On("GetProducts").Return(nil, expectedErr)

	actualProducts, actualErr := productService.GetProducts()

	assert.Nil(t, actualProducts, "products should be nil")
	assert.NotNil(t, actualErr, "error should not be nil")
	assert.Equal(t, expectedErr, actualErr, "error should be equal")
}

func TestProductService_GetOneProduct_Found(t *testing.T) {
	expectedProduct := &entity.Product{
		GormModel: entity.GormModel{
			ID: 3,
		},
		Title:       "Product 1",
		Description: "Product 1 Description",
	}
	productRepository.Mock.On("GetProductById", 3).Return(expectedProduct, nil)

	actualProduct, actualErr := productService.GetProductById(3)

	assert.Nil(t, actualErr, "error should be nil")
	assert.NotNil(t, actualProduct, "product should not be nil")
	assert.Equal(t, expectedProduct.GormModel.ID, actualProduct.ID, "product id should be equal")
	assert.Equal(t, expectedProduct.Title, actualProduct.Title, "product title should be equal")
	assert.Equal(t, expectedProduct.Description, actualProduct.Description, "product description should be equal")
}

func TestProductService_GetOneProduct_NotFound(t *testing.T) {
	expectedErr := errors.NewNotFoundError("Product not found")
	productRepository.Mock.On("GetProductById", -1).Return(nil, expectedErr)

	actualProduct, actualErr := productService.GetProductById(-1)

	assert.Nil(t, actualProduct, "product should be nil")
	assert.NotNil(t, actualErr, "error should not be nil")
	assert.Equal(t, expectedErr, actualErr, "error should be equal")
}
