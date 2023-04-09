package repo_interfaces

import (
	"go-product/domain/entity"
	"go-product/pkg/errors"
)

type ProductRepository interface {
	CreateProduct(product *entity.Product) (*entity.Product, errors.Error)
	UpdateProduct(id int, product *entity.Product) (*entity.Product, errors.Error)
	DeleteProduct(id int) errors.Error
	GetProducts() ([]*entity.Product, errors.Error)
	GetProductById(id int) (*entity.Product, errors.Error)
}
