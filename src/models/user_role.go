package models

type UserRole struct {
	Base
	UserID int64 `json:"user_id"`
	RoleID int64 `json:"role_id"`
}
