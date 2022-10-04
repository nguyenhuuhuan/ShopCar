package models

type Role struct {
	Base
	RoleName string `json:"role_name"`
	Status   string `json:"status"`
	Code     string `json:"code"`
}
