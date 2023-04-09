package dto

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
