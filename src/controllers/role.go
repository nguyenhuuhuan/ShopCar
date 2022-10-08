package controllers

import (
	"Improve/src/dtos"
	"Improve/src/errors"
	"Improve/src/services"
	"Improve/src/utils"
	"github.com/gin-gonic/gin"
)

type RoleController interface {
	Create(ctx *gin.Context)
}

type roleController struct {
	Base
	roleService services.RoleService
}

// Create role
// @Title Create role
// @Description Create role
// @Tags ShopCar
// @Param body body dtos.CreateRoleRequest true "data"
// @Success 200 {object} dtos.CreateRoleResponse
// @Failure 400 {object} errors.AppError
// @Failure 403 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Security JWTAccessToken
// @router /role [post]
func (r roleController) Create(ctx *gin.Context) {
	var role *dtos.CreateRoleRequest
	err := ctx.ShouldBindJSON(&role)
	if err != nil {
		utils.HandleError(ctx, errors.New(errors.InvalidRequestError))
		return
	}
	resp, err := r.roleService.Create(ctx.Request.Context(), role)
	if err != nil {
		r.HandleError(ctx, err)
		return
	}
	r.JSON(ctx, resp)
}

func NewRoleController(roleService services.RoleService) RoleController {
	return &roleController{
		roleService: roleService,
	}

}
