package entity

import "go-product/domain/dto"

type User struct {
	GormModel
	FullName string `gorm:"notNull"`
	Email    string `gorm:"notNull;unique"`
	Password string `gorm:"notNull"`
	IsAdmin  bool   `gorm:"default:false"`
	Products []Product
}

func (u *User) ToResponse() dto.UserResponse {
	var products []dto.ProductResponse
	for _, p := range u.Products {
		products = append(products, p.ToResponse())
	}

	return dto.UserResponse{
		ID:       u.ID,
		FullName: u.FullName,
		Email:    u.Email,
		Products: products,
	}
}

func (u *User) FromRequest(req dto.UserRequest) {
	u.FullName = req.FullName
	u.Email = req.Email
	u.Password = req.Password
}
