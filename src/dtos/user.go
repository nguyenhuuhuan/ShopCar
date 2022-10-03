package dtos

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterRequest struct {
	UserName string `json:"user_name" binding:"omitempty"`
	Email    string `json:"email" binding:"required"`
	Provider string `json:"provider" binding:"required"`
	Status   string `json:"status" binding:"oneof=ACTIVE INACTIVE"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name" binding:"omitempty"`
}

type UserRegisterResponse struct {
	Meta Meta                 `json:"meta"`
	Data *UserRegisterRequest `json:"data"`
}
