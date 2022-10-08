package controllers

import (
	"Improve/src/dtos"
	"Improve/src/errors"
	"Improve/src/logger"
	"Improve/src/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController interface {
	List(ctx *gin.Context)
	GetUser(ctx *gin.Context)
}
type userController struct {
	Base
	userService services.UserService
}

func (u userController) List(ctx *gin.Context) {
	var listUserReq dtos.ListUserRequest
	err := ctx.ShouldBindQuery(&listUserReq)
	if err != nil {
		logger.Context(ctx).Errorf("[UserController][GetUser] Invalid data %v: ", err)
		u.HandleError(ctx, errors.New(errors.InvalidRequestError, err))
		return
	}
	resp, err := u.userService.List(ctx, &listUserReq)
	u.Respond(ctx, resp, err)
}

func (u userController) GetUser(ctx *gin.Context) {
	idx := ctx.Param("id")
	id, err := strconv.ParseInt(idx, 10, 64)
	if err != nil {
		logger.Context(ctx).Errorf("[UserController][GetUser] Invalid data %v: ", err)
		u.HandleError(ctx, errors.New(errors.InvalidRequestError, err))
		return
	}
	resp, err := u.userService.GetUser(ctx, id)
	u.Respond(ctx, resp, err)
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}
