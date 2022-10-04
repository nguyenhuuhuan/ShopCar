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

func (r roleController) Create(ctx *gin.Context) {
	var role *dtos.Role
	err := ctx.ShouldBindJSON(role)
	if err != nil {
		utils.HandleError(ctx, errors.New(errors.InvalidRequestError))
		return
	}
	resp, err := r.roleService.Create(ctx.Request.Context(), role)
	r.Respond(ctx, resp, err)
}

func NewRoleController(roleService services.RoleService) RoleController {
	return &roleController{
		roleService: roleService,
	}

}
