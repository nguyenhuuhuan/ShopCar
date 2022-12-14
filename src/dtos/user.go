package dtos

type UserLoginRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterRequest struct {
	UserName    string `json:"user_name" binding:"omitempty"`
	Email       string `json:"email" binding:"required" validate:"required"`
	Owner       string `json:"owner"`
	Provider    string `json:"provider" binding:"required"`
	Status      string `json:"status" binding:"oneof=ACTIVE INACTIVE"`
	Password    string `json:"password" binding:"required" validate:"required"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number" binding:"required" validate:"required"`
	Roles       []Role `json:"role"`
}

type User struct {
	Base
	UserName    string `json:"user_name" binding:"omitempty"`
	Email       string `json:"email" binding:"required"`
	Owner       string `json:"owner"`
	Provider    string `json:"provider" binding:"required"`
	Status      string `json:"status" binding:"oneof=ACTIVE INACTIVE"`
	Password    string `json:"password" binding:"required"`
	FullName    string `json:"full_name" binding:"omitempty"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type UserRegisterResponse struct {
	Meta Meta     `json:"meta"`
	Data UserInfo `json:"data"`
}

type UserInfo struct {
	User        User   `json:"user"`
	AccessToken string `json:"access_token"`
}
type UserLoginResponse struct {
	Meta Meta     `json:"meta"`
	Data UserInfo `json:"data"`
}

type ListUserRequest struct {
	Base
	Username    string `form:"username"`
	Email       string `form:"email"`
	Owner       string `form:"owner"`
	PhoneNumber string `form:"phone_number"`
	Status      string `form:"status"`
	SortBy      string `form:"sort_by" binding:"omitempty,oneof=username email id owner"`
	Page        int    `form:"page" binding:"omitempty,gte=0"`
	PageSize    int    `form:"page_size" binding:"required,gte=0,max=100"`
	Reverse     bool   `form:"reverse"`
}
type ListUserResponse struct {
	Meta PaginationMeta `json:"meta"`
	Data []User         `json:"data"`
}

type GetUserResponse struct {
	Meta Meta `json:"meta"`
	Data User `json:"data"`
}
