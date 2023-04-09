package repositories

import (
	"go-product/domain/entity"
	"go-product/pkg/errors"
	"go-product/repositories/i_repositories"
	"gorm.io/gorm"
	"log"
)

type productRepository struct {
	*gorm.DB
}

func NewProductRepository(db *gorm.DB) i_repositories.ProductRepository {
	return &productRepository{db}
}

func (p productRepository) CreateProduct(product *entity.Product) (*entity.Product, errors.Error) {
	err := p.DB.Model(&entity.Product{}).Create(product).Error
	if err != nil {
		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return product, nil
}

func (p productRepository) UpdateProduct(id int, product *entity.Product) (*entity.Product, errors.Error) {
	var updatedProduct = &entity.Product{}

	err := p.DB.Model(updatedProduct).
		Where("id = ?", id).
		Updates(&entity.Product{
			Title:       product.Title,
			Description: product.Description,
		}).
		First(updatedProduct).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errs := errors.NewNotFoundError("The product you are trying to update does not exist")
			return nil, errs
		}

		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return updatedProduct, nil
}

func (p productRepository) DeleteProduct(id int) errors.Error {
	var product = &entity.Product{}
	err := p.DB.Model(&entity.Product{}).Where("id = ?", id).First(product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errs := errors.NewNotFoundError("The product you are trying to delete does not exist")
			return errs
		}
	}

	err = p.DB.Model(&entity.Product{}).Delete("id = ?", id).Error
	if err != nil {
		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return errs
	}

	return nil
}

func (p productRepository) GetProducts() ([]*entity.Product, errors.Error) {
	var products []*entity.Product

	err := p.DB.Model(&entity.Product{}).Find(&products).Error
	if err != nil {
		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	if len(products) == 0 {
		products = []*entity.Product{}
	}

	return products, nil
}

func (p productRepository) GetProductById(id int) (*entity.Product, errors.Error) {
	var product = &entity.Product{}

	err := p.DB.
		Model(&entity.Product{}).
		Where("id = ?", id).
		First(product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errs := errors.NewNotFoundError("The product you are trying to retrieve does not exist")
			return nil, errs
		}

		log.Println(err)
		errs := errors.NewInternalServerError("An error occurred while processing your request. Please try again later")
		return nil, errs
	}

	return product, nil
}
