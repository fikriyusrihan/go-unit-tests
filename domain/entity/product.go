package entity

import "go-product/domain/dto"

type Product struct {
	GormModel
	Title       string `gorm:"notNull"`
	Description string `gorm:"notNull"`
	UserID      uint
}

func (p *Product) ToResponse() dto.ProductResponse {
	return dto.ProductResponse{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
	}
}

func (p *Product) FromRequest(req dto.ProductRequest) {
	p.Title = req.Title
	p.Description = req.Description
}
