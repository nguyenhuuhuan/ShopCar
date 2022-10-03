package controllers

import (
	"Improve/src/dtos"
	"Improve/src/errors"
	"Improve/src/services"
	"Improve/src/utils"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authController struct {
	Base
	authService services.AuthService
}

func (a *authController) Register(ctx *gin.Context) {
	var registerReq dtos.UserRegisterRequest
	err := ctx.ShouldBindJSON(&registerReq)
	if err != nil {
		utils.HandleError(ctx, errors.New(errors.InvalidRequestError))
		return
	}
	resp, err := a.authService.Register(ctx.Request.Context(), &registerReq)
	a.Respond(ctx, resp, err)
}

func (a authController) Login(ctx *gin.Context) {

}

func NewAuthController(authService services.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}
