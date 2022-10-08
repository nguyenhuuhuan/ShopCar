package repositories

import (
	"Improve/src/errors"
	"Improve/src/models"
	"context"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id int64) (*models.User, error)
	IsDuplicateEmail(ctx context.Context, email, username string) bool
	GetByEmailOrUser(ctx context.Context, email, username string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
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
		log.Printf("[UserRepo] Create user is failed %v: ", err)
		return nil, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
