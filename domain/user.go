package domain

type User struct {
	GormModel
	FullName string `gorm:"notNull"`
	Email    string `gorm:"notNull;unique"`
	Password string `gorm:"notNull"`
	Products []Product
}

type UserRequest struct {
	FullName string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserResponse struct {
	ID       uint              `json:"id"`
	FullName string            `json:"fullName"`
	Email    string            `json:"email"`
	Products []ProductResponse `json:"products"`
}

func (u *User) ToResponse() UserResponse {
	var products []ProductResponse
	for _, p := range u.Products {
		products = append(products, p.ToResponse())
	}

	return UserResponse{
		ID:       u.ID,
		FullName: u.FullName,
		Email:    u.Email,
		Products: products,
	}
}

func (u *User) FromRequest(req UserRequest) {
	u.FullName = req.FullName
	u.Email = req.Email
	u.Password = req.Password
}
