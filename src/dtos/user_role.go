package dtos

type UserRole struct {
	Base
	UserID int64 `json:"user_id" validate:"required"`
	RoleID int64 `json:"role_id" validate:"required"`
}

type UserRoleResponse struct {
	Meta *Meta     `json:"meta"`
	Data *UserRole `json:"data"`
}
