package repository

import (
	"go-product/domain"
	"gorm.io/gorm"
)

type productRepository struct {
	*gorm.DB
}

type ProductRepository interface {
	CreateProduct(product *domain.Product) (*domain.Product, error)
	UpdateProduct(id uint, product *domain.Product) (*domain.Product, error)
	DeleteProduct(id uint) error
	GetProducts() ([]*domain.Product, error)
	GetProductById(id uint) (*domain.Product, error)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (p productRepository) CreateProduct(product *domain.Product) (*domain.Product, error) {
	err := p.DB.Model(&domain.Product{}).Create(product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p productRepository) UpdateProduct(id uint, product *domain.Product) (*domain.Product, error) {
	var updatedProduct = &domain.Product{}

	err := p.DB.Model(updatedProduct).
		Where("id = ?", id).
		Updates(&domain.Product{
			Title:       product.Title,
			Description: product.Description,
		}).
		First(updatedProduct).Error

	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (p productRepository) DeleteProduct(id uint) error {
	return p.DB.Model(&domain.Product{}).Delete("id = ?", id).Error
}

func (p productRepository) GetProducts() ([]*domain.Product, error) {
	var products []*domain.Product

	err := p.DB.Model(&domain.Product{}).Find(&products).Error
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		products = []*domain.Product{}
	}

	return products, nil
}

func (p productRepository) GetProductById(id uint) (*domain.Product, error) {
	var product = &domain.Product{}

	err := p.DB.
		Model(&domain.Product{}).
		Where("id = ?", id).
		First(product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}
