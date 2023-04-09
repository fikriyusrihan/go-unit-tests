package i_repositories

import "go-product/domain/entity"

type ProductRepository interface {
	CreateProduct(product *entity.Product) (*entity.Product, error)
	UpdateProduct(id int, product *entity.Product) (*entity.Product, error)
	DeleteProduct(id int) error
	GetProducts() ([]*entity.Product, error)
	GetProductById(id int) (*entity.Product, error)
}
