package repositories

import (
	"Improve/src/dtos"
	"Improve/src/errors"
	"Improve/src/logger"
	"Improve/src/models"
	"context"
	"fmt"
	"gorm.io/gorm"
)

const (
	SortByEmail     = "email"
	SortByID        = "id"
	SortByUsername  = "username"
	SortByCreatedAt = "created_at"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id int64) (*models.User, error)
	List(ctx context.Context, req *dtos.ListUserRequest) ([]models.User, int64, error)
	IsDuplicateEmail(ctx context.Context, email, username string) bool
	GetByEmailOrUser(ctx context.Context, email, username string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) List(ctx context.Context, req *dtos.ListUserRequest) ([]models.User, int64, error) {
	var (
		listUser      []models.User
		order         = models.DESC
		count         = int64(0)
		offset, limit int
		err           error
		query         = u.db.WithContext(ctx)
	)
	if req.Page == 0 && req.PageSize == 0 {
		limit, offset = -1, -1
	} else {
		limit = req.PageSize
		offset = (req.Page - 1) * req.PageSize
	}

	if req.ID != 0 {
		query = query.Where("users.id = ? ", req.ID)
	}

	if req.Username != "" {
		query = query.Where("users.username LIKE ? ", fmt.Sprintf("%%%v%%", req.Username))
	}

	if req.Status != "" {
		query = query.Where("users.status = ? ", req.Status)
	}

	if req.Email != "" {
		query = query.Where("users.email LIKE ? ", req.Email)
	}
	if req.Reverse {
		order = models.ASC
	}
	switch req.SortBy {
	case SortByEmail:
		query = query.Order(fmt.Sprintf("%v %v", SortByEmail, order))
	case SortByID:
		query = query.Order(fmt.Sprintf("%v %v", SortByID, order))
	case SortByUsername:
		query = query.Order(fmt.Sprintf("%v %v", SortByUsername, order))
	default:
		query = query.Order(fmt.Sprintf("%v %v", SortByCreatedAt, order))
	}

	err = query.Model(models.User{}).Count(&count).
		Where("owner = ?", req.Owner).
		Limit(limit).
		Offset(offset).
		Find(&listUser).Error

	if err != nil {
		return nil, 0, err
	}

	return listUser, count, nil

}

func (u userRepository) GetByEmailOrUser(ctx context.Context, email, username string) (*models.User, error) {
	var user *models.User
	if email == "" {
		err := u.db.WithContext(ctx).Where("username = ?", username).Take(&user).Error
		if err != nil {
			return user, err
		}
		return user, nil
	} else if username == "" {
		err := u.db.WithContext(ctx).Where("email = ?", email).Take(&user).Error
		if err != nil {
			return user, err
		}
		return user, nil
	}
	return nil, errors.New(errors.InternalServerError)
}

func (u userRepository) IsDuplicateEmail(ctx context.Context, email, username string) bool {
	var user models.User
	if username == "" {
		err := u.db.WithContext(ctx).Where("email = ?", email).Take(&user).Error
		if err != nil {
			return false
		}
	} else if email == "" {
		err := u.db.WithContext(ctx).Where("username = ?", username).Take(&user).Error
		if err != nil {
			return false
		}
	} else {
		return true
	}
	return true
}

func (u userRepository) Create(ctx context.Context, user *models.User) error {
	return u.db.WithContext(ctx).Create(&user).Error
}

func (u userRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	var user *models.User
	err := u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Context(ctx).Errorf("[UserRepo] Create user is failed %v: ", err)
		return nil, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
