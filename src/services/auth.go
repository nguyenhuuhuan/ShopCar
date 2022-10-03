package services

import (
	"Improve/src/dtos"
	"Improve/src/errors"
	"Improve/src/repositories"
	"context"
	"fmt"
	_ "github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/copier"
	"net/http"
	_ "time"
)

type AuthService interface {
	Register(ctx context.Context, req *dtos.UserRegisterRequest) (*dtos.UserRegisterResponse, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func (a *authService) Register(ctx context.Context, req *dtos.UserRegisterRequest) (*dtos.UserRegisterResponse, error) {
	var data *dtos.UserRegisterRequest
	_ = copier.Copy(&req, data)

	if checkIsDuplicateEmail := a.userRepo.IsDuplicateEmail(ctx, data.Email, data.UserName); checkIsDuplicateEmail {
		fmt.Errorf("[AuthService][Register] Email is duplicated")
		return nil, errors.New(errors.DuplicateError)
	}

	//maker, err := token.NewJWTMaker(configs.secretKeyMaker)

	return &dtos.UserRegisterResponse{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: "OK",
		},
		Data: data,
	}, nil

}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}
