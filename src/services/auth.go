package services

import (
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
	"strconv"
	"time"
	_ "time"
)

type AuthService interface {
	Register(ctx context.Context, req *dtos.UserRegisterRequest) (*dtos.UserRegisterResponse, error)
}

type authService struct {
	userRepo     repositories.UserRepository
	roleRepo     repositories.RoleRepository
	userRoleRepo repositories.UserRoleRepository
}

func (a *authService) Register(ctx context.Context, req *dtos.UserRegisterRequest) (*dtos.UserRegisterResponse, error) {
	var (
		user *models.User
		data *dtos.User
	)
	err := utils.ValidateData(req)
	if err != nil {
		fmt.Errorf("[AuthService][Register] User is invalid")
		return nil, errors.New(errors.UnsupportedEntityError)
	}

	if checkIsDuplicateEmail := a.userRepo.IsDuplicateEmail(ctx, req.Email, req.UserName); checkIsDuplicateEmail {
		fmt.Errorf("[AuthService][Register] Email is duplicated")
		return nil, errors.New(errors.DuplicateError)
	}

	_ = copier.Copy(&user, req)
	userID, err := a.userRepo.Create(ctx, *user)
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
			UserID: userID,
			RoleID: role.ID,
		}

		err = a.userRoleRepo.Create(ctx, &userRole)
		if err != nil {
			fmt.Errorf("[AuthService][Register] Err Create UserRole %v ", err)
			return nil, errors.New(errors.InternalServerError)
		}
	}

	maker, err := token.NewJWTMaker(strconv.FormatUint(uint64(user.ID), 10))
	if err != nil {
		fmt.Errorf("[AuthService][Register] Init Maker error %v ", err)
		return nil, errors.New(errors.InternalServerError)
	}
	token, err := maker.CreateToken(user.Email, 2*time.Hour)
	if err != nil {
		fmt.Errorf("[AuthService][Register] Create token error %v ", err)
		return nil, errors.New(errors.InternalServerError)
	}

	data.Token = token
	_ = copier.Copy(&data, user)
	return &dtos.UserRegisterResponse{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: "OK",
		},
		Data: data,
	}, nil

}

func NewAuthService(userRepo repositories.UserRepository, roleRepo repositories.RoleRepository, userRoleRepo repositories.UserRoleRepository) AuthService {
	return &authService{
		userRepo:     userRepo,
		roleRepo:     roleRepo,
		userRoleRepo: userRoleRepo,
	}
}
