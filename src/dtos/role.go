package dtos

type Role struct {
	Base
	RoleName string `json:"role_name"`
	Status   string `json:"status" binding:"oneof=ACTIVE INACTIVE"`
	Code     string `json:"code"`
}

type CreateRoleRequest struct {
	RoleName string `json:"role_name" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Status   string `json:"status" binding:"required,oneof=ACTIVE INACTIVE"`
}
type GetRoleResponse struct {
	Meta *Meta `json:"meta"`
	Data *Role `json:"data"`
}
type CreateRoleResponse struct {
	Meta *Meta `json:"meta"`
	Data *Role `json:"data"`
}
