package dtos

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterRequest struct {
	UserName    string `json:"user_name" binding:"omitempty"`
	Email       string `json:"email" binding:"required" validate:"required,email"`
	Provider    string `json:"provider" binding:"required"`
	Status      string `json:"status" binding:"oneof=ACTIVE INACTIVE"`
	Password    string `json:"password" binding:"required" validate:"required,passwd"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number" binding:"required" validate:"required"`
	Roles       []Role `json:"role"`
}

type User struct {
	Base
	UserName    string `json:"user_name" binding:"omitempty"`
	Email       string `json:"email" binding:"required"`
	Provider    string `json:"provider" binding:"required"`
	Status      string `json:"status" binding:"oneof=ACTIVE INACTIVE"`
	Password    string `json:"password" binding:"required"`
	FullName    string `json:"full_name" binding:"omitempty"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Token       string `json:"token"`
}

type UserRegisterResponse struct {
	Meta Meta  `json:"meta"`
	Data *User `json:"data"`
}
