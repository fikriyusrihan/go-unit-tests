package repositories_mock

import (
	"github.com/stretchr/testify/mock"
	"go-product/domain/entity"
	"go-product/pkg/errors"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (r *ProductRepositoryMock) CreateProduct(product *entity.Product) (*entity.Product, errors.Error) {
	args := r.Mock.Called(product)
	if args.Get(0) == nil {
		return nil, args.Get(1).(errors.Error)
	}

	return args.Get(0).(*entity.Product), nil
}

func (r *ProductRepositoryMock) UpdateProduct(id int, product *entity.Product) (*entity.Product, errors.Error) {
	args := r.Mock.Called(id, product)
	if args.Get(0) == nil {
		return nil, args.Get(1).(errors.Error)
	}

	return args.Get(0).(*entity.Product), nil
}

func (r *ProductRepositoryMock) DeleteProduct(id int) errors.Error {
	args := r.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(errors.Error)
}

func (r *ProductRepositoryMock) GetProducts() ([]*entity.Product, errors.Error) {
	args := r.Mock.Called()
	if args.Get(0) == nil {
		return nil, args.Get(1).(errors.Error)
	}

	return args.Get(0).([]*entity.Product), nil
}

func (r *ProductRepositoryMock) GetProductById(id int) (*entity.Product, errors.Error) {
	args := r.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Get(1).(errors.Error)
	}

	return args.Get(0).(*entity.Product), nil
}
