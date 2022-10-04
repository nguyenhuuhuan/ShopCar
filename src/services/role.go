package services

import (
	"Improve/src/dtos"
	"Improve/src/errors"
	"Improve/src/models"
	"Improve/src/repositories"
	"Improve/src/utils"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"net/http"
)

type RoleService interface {
	Create(ctx context.Context, role *dtos.Role) (*dtos.CreateRoleResponse, error)
	GetByID(ctx context.Context, id int64) (*dtos.GetRoleResponse, error)
}

type roleService struct {
	roleRepo repositories.RoleRepository
}

func (r roleService) Create(ctx context.Context, data *dtos.Role) (*dtos.CreateRoleResponse, error) {
	var role models.Role
	err := utils.ValidateData(role)
	if err != nil {
		fmt.Errorf("[RoleService][Create] Role is invalid")
		return nil, errors.New(errors.UnsupportedEntityError)
	}

	_ = copier.Copy(&role, data)
	err = r.roleRepo.Create(ctx, &role)
	if err != nil {
		fmt.Errorf("[RoleService][Create] error Create Role %v ", err)
		return nil, errors.New(errors.InternalServerError)
	}

	return &dtos.CreateRoleResponse{
		Meta: &dtos.Meta{
			Code:    http.StatusOK,
			Message: "OK",
		},
		Data: data,
	}, nil
}

func (r roleService) GetByID(ctx context.Context, id int64) (*dtos.GetRoleResponse, error) {
	var (
		role *models.Role
		data *dtos.Role
	)
	role, err := r.roleRepo.GetByID(ctx, id)
	if err != nil {
		fmt.Errorf("[RoleService][GetByID] Role not found %v: ", err)
		return nil, err
	}

	_ = copier.Copy(&data, role)
	return &dtos.GetRoleResponse{
		Meta: &dtos.Meta{
			Code:    http.StatusOK,
			Message: "OK",
		},
		Data: data,
	}, nil
}

func NewRoleService(roleRepo repositories.RoleRepository) RoleService {
	return &roleService{
		roleRepo: roleRepo,
	}
}
