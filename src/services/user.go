package services

import (
	"Improve/src/dtos"
	"Improve/src/errors"
	"Improve/src/logger"
	"Improve/src/repositories"
	"context"
	"github.com/jinzhu/copier"
	"net/http"
)

type UserService interface {
	List(ctx context.Context, req *dtos.ListUserRequest) (*dtos.ListUserResponse, error)
	GetUser(ctx context.Context, id int64) (*dtos.GetUserResponse, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func (u *userService) List(ctx context.Context, req *dtos.ListUserRequest) (*dtos.ListUserResponse, error) {
	var (
		data  []dtos.User
		count int64
	)

	listUser, count, err := u.userRepo.List(ctx, req)
	if err != nil {
		logger.Context(ctx).Errorf("[UserService][List] List User is not found %v", err)
		return nil, errors.New(errors.InternalServerError)
	}
	_ = copier.Copy(&data, listUser)
	return &dtos.ListUserResponse{
		Meta: dtos.PaginationMeta{
			Meta: dtos.Meta{
				Code:    http.StatusOK,
				Message: "OK",
			},
			Total:    count,
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Data: data,
	}, nil
}

func (u *userService) GetUser(ctx context.Context, id int64) (*dtos.GetUserResponse, error) {
	var data dtos.User
	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		logger.Context(ctx).Errorf("[UserService][GetUser] User is not found %v", err)
		return nil, errors.New(errors.InternalServerError)
	}
	_ = copier.Copy(&data, user)
	return &dtos.GetUserResponse{
		Meta: dtos.Meta{
			Code:    http.StatusOK,
			Message: "OK",
		},
		Data: data,
	}, nil
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
