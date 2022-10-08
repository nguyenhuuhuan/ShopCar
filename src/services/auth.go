package services

import (
	"Improve/src/configs"
	"Improve/src/dtos"
	"Improve/src/errors"
	"Improve/src/models"
	"Improve/src/repositories"
	"Improve/src/token"
	"Improve/src/utils"
	"context"
	"fmt"
	_ "github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/copier"
	"net/http"
	_ "time"
)

const UserActive = "ACTIVE"

type AuthService interface {
	Register(ctx context.Context, req *dtos.UserRegisterRequest) (*dtos.UserRegisterResponse, error)
	Login(ctx context.Context, req *dtos.UserLoginRequest) (*dtos.UserLoginResponse, error)
}

type authService struct {
	app          *configs.App
	tokenMaker   token.Maker
	userRepo     repositories.UserRepository
	roleRepo     repositories.RoleRepository
	userRoleRepo repositories.UserRoleRepository
}

func (a *authService) Login(ctx context.Context, req *dtos.UserLoginRequest) (*dtos.UserLoginResponse, error) {
	var (
		userInfo dtos.UserInfo
	)
	user, err := a.userRepo.GetByEmailOrUser(ctx, req.Email, req.Username)
	if err != nil {
		fmt.Printf("[AuthService][Register] Username or Email is invalid %v", err)
		return nil, errors.New(errors.InternalServerError)
	}

	if checkPassword := user.ComparePassword(req.Password, user.Password); !checkPassword {
		fmt.Printf("[AuthService][Register] Password is invalid %v", err)
		return nil, errors.New(errors.PasswordInvalid)
	}

	accessToken, err := a.tokenMaker.CreateToken(req.Email, a.app.JWT.AccessTokenDuration)
	if err != nil {
		fmt.Printf("[AuthService][Register] Token is invalid %v", err)
		return nil, errors.New(errors.InternalServerError)
	}

	_ = copier.Copy(&userInfo.User, user)
	userInfo.AccessToken = accessToken
	return &dtos.UserLoginResponse{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: "OK",
		},
		Data: userInfo,
	}, nil

}

func (a *authService) Register(ctx context.Context, req *dtos.UserRegisterRequest) (*dtos.UserRegisterResponse, error) {
	var (
		user models.User
		data dtos.UserInfo
	)
	err := utils.ValidateData(req)
	if err != nil {
		fmt.Printf("[AuthService][Register] User is invalid %v", err)
		return nil, errors.New(errors.UnsupportedEntityError)
	}

	if checkIsDuplicateEmail := a.userRepo.IsDuplicateEmail(ctx, req.Email, req.UserName); checkIsDuplicateEmail {
		fmt.Errorf("[AuthService][Register] Email is duplicated")
		return nil, errors.New(errors.DuplicateError)
	}

	_ = copier.Copy(&user, req)
	user.Password, err = user.HashPassword(user.Password)
	if err != nil {
		fmt.Errorf("[AuthService][Register] Password hash failed %v", err)
		return nil, errors.New(errors.UnsupportedEntityError)
	}
	err = a.userRepo.Create(ctx, &user)
	if err != nil {
		fmt.Errorf("[AuthService][Register] error Create User %v ", err)
		return nil, errors.New(errors.InternalServerError)
	}

	for _, val := range req.Roles {
		role, err := a.roleRepo.GetByID(ctx, val.ID)
		if err != nil {
			fmt.Errorf("[AuthService][Register] Role not found")
			return nil, errors.New(errors.InternalServerError)
		}
		userRole := models.UserRole{
			UserID: user.ID,
			RoleID: role.ID,
		}

		err = a.userRoleRepo.Create(ctx, &userRole)
		if err != nil {
			fmt.Errorf("[AuthService][Register] Err Create UserRole %v ", err)
			return nil, errors.New(errors.InternalServerError)
		}
	}

	token, err := a.tokenMaker.CreateToken(user.Email, a.app.JWT.AccessTokenDuration)
	if err != nil {
		fmt.Errorf("[AuthService][Register] Create token error %v ", err)
		return nil, errors.New(errors.InternalServerError)
	}

	data.AccessToken = token
	_ = copier.Copy(&data.User, user)
	return &dtos.UserRegisterResponse{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: "OK",
		},
		Data: data,
	}, nil

}

func NewAuthService(app configs.App, tokenMake token.Maker, userRepo repositories.UserRepository, roleRepo repositories.RoleRepository, userRoleRepo repositories.UserRoleRepository) AuthService {
	return &authService{
		app:          &app,
		tokenMaker:   tokenMake,
		userRepo:     userRepo,
		roleRepo:     roleRepo,
		userRoleRepo: userRoleRepo,
	}
}
