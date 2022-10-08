package services

import (
	"Improve/src/dtos"
	"Improve/src/errors"
	"Improve/src/logger"
	"Improve/src/models"
	"Improve/src/repositories"
	"Improve/src/utils"
	"context"
	"github.com/jinzhu/copier"
	"net/http"
)

type RoleService interface {
	Create(ctx context.Context, role *dtos.CreateRoleRequest) (*dtos.CreateRoleResponse, error)
	GetByID(ctx context.Context, id int64) (*dtos.GetRoleResponse, error)
}

type roleService struct {
	roleRepo repositories.RoleRepository
}

func (r roleService) Create(ctx context.Context, req *dtos.CreateRoleRequest) (*dtos.CreateRoleResponse, error) {
	var (
		role models.Role
		data dtos.Role
	)
	err := utils.ValidateData(role)
	if err != nil {
		logger.Context(ctx).Errorf("[RoleService][Create] Role is invalid %v: ", err)
		return nil, errors.New(errors.UnsupportedEntityError)
	}

	if checkExistRole := r.roleRepo.IsExistRoleAndCode(ctx, req.RoleName, req.Code); checkExistRole {
		logger.Context(ctx).Errorf("[RoleService][Create] Role or Code is exists")
		return nil, errors.New(errors.RoleIsExistedError)

	}
	_ = copier.Copy(&role, req)
	err = r.roleRepo.Create(ctx, &role)
	if err != nil {
		logger.Context(ctx).Errorf("[RoleService][Create] error Create Role %v: ", err)
		return nil, errors.New(errors.InternalServerError)
	}

	_ = copier.Copy(&data, role)
	return &dtos.CreateRoleResponse{
		Meta: &dtos.Meta{
			Code:    http.StatusOK,
			Message: "OK",
		},
		Data: &data,
	}, nil
}

func (r roleService) GetByID(ctx context.Context, id int64) (*dtos.GetRoleResponse, error) {
	var (
		role *models.Role
		data *dtos.Role
	)
	role, err := r.roleRepo.GetByID(ctx, id)
	if err != nil {
		logger.Context(ctx).Errorf("[RoleService][GetByID] Role not found %v: ", err)
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
