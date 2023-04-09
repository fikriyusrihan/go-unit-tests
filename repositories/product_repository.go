package repositories

import (
	"go-product/domain/entity"
	"go-product/repositories/i_repositories"
	"gorm.io/gorm"
)

type productRepository struct {
	*gorm.DB
}

func NewProductRepository(db *gorm.DB) i_repositories.ProductRepository {
	return &productRepository{db}
}

func (p productRepository) CreateProduct(product *entity.Product) (*entity.Product, error) {
	err := p.DB.Model(&entity.Product{}).Create(product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p productRepository) UpdateProduct(id int, product *entity.Product) (*entity.Product, error) {
	var updatedProduct = &entity.Product{}

	err := p.DB.Model(updatedProduct).
		Where("id = ?", id).
		Updates(&entity.Product{
			Title:       product.Title,
			Description: product.Description,
		}).
		First(updatedProduct).Error

	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (p productRepository) DeleteProduct(id int) error {
	var product = &entity.Product{}
	err := p.DB.Model(&entity.Product{}).Where("id = ?", id).First(product).Error
	if err != nil {
		return err
	}

	return p.DB.Model(&entity.Product{}).Delete("id = ?", id).Error
}

func (p productRepository) GetProducts() ([]*entity.Product, error) {
	var products []*entity.Product

	err := p.DB.Model(&entity.Product{}).Find(&products).Error
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		products = []*entity.Product{}
	}

	return products, nil
}

func (p productRepository) GetProductById(id int) (*entity.Product, error) {
	var product = &entity.Product{}

	err := p.DB.
		Model(&entity.Product{}).
		Where("id = ?", id).
		First(product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}
