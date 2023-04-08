package domain

type Product struct {
	GormModel
	Title       string `gorm:"notNull"`
	Description string `gorm:"notNull"`
	UserID      uint
}

type ProductRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type ProductResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (p *Product) ToResponse() ProductResponse {
	return ProductResponse{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
	}
}

func (p *Product) FromRequest(req ProductRequest) {
	p.Title = req.Title
	p.Description = req.Description
}
