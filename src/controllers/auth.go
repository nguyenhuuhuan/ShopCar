package controllers

import (
	"Improve/src/dtos"
	"Improve/src/errors"
	"Improve/src/logger"
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

// Register for user
// @Title Register user
// @Description Register user
// @Tags ShopCar
// @Param body body dtos.UserRegisterRequest true "data"
// @Success 200 {object} dtos.UserRegisterResponse
// @Failure 400 {object} errors.AppError
// @Failure 403 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Security JWTAccessToken
// @router /auth/register [post]
func (a *authController) Register(ctx *gin.Context) {
	var registerReq dtos.UserRegisterRequest
	err := ctx.ShouldBindJSON(&registerReq)
	if err != nil {
		logger.Context(ctx).Errorf("[AuthController][Register] Error validate %v", err)
		utils.HandleError(ctx, errors.New(errors.InvalidRequestError))
		return
	}
	resp, err := a.authService.Register(ctx.Request.Context(), &registerReq)
	a.Respond(ctx, resp, err)
}

// Login user
// @Title user Login
// @Description user Login
// @Tags ShopCar
// @Param body body dtos.UserLoginRequest true "data"
// @Success 200 {object} dtos.UserLoginResponse
// @Failure 400 {object} errors.AppError
// @Failure 403 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Security JWTAccessToken
// @router /auth/login [post]
func (a authController) Login(ctx *gin.Context) {
	var loginReq dtos.UserLoginRequest
	err := ctx.ShouldBindJSON(&loginReq)
	if err != nil {
		logger.Context(ctx).Errorf("[AuthController][Register] Error validate %v", err)
		utils.HandleError(ctx, errors.New(errors.InvalidRequestError))
		return
	}
	resp, err := a.authService.Login(ctx.Request.Context(), &loginReq)
	a.Respond(ctx, resp, err)
}

func NewAuthController(authService services.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}
