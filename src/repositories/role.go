package repositories

import (
	"Improve/src/models"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) error
	GetByID(ctx context.Context, id int64) (*models.User, error)
	IsDuplicateEmail(ctx context.Context, email, username string) bool
}

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) IsDuplicateEmail(ctx context.Context, email, username string) bool {
	err := u.db.WithContext(ctx).Where("email = ? or username = ?", email, username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return err == nil
}

func (u userRepository) Create(ctx context.Context, user models.User) error {
	err := u.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		log.Printf("[UserRepo] Create user is failed %v: ", err)
		return err
	}
	return nil
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
