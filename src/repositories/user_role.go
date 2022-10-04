package repositories

import (
	"Improve/src/models"
	"context"
	"gorm.io/gorm"
	"log"
)

type UserRoleRepository interface {
	Create(ctx context.Context, role *models.UserRole) error
}

type userRoleRepository struct {
	db *gorm.DB
}

func (u userRoleRepository) Create(ctx context.Context, userRole *models.UserRole) error {
	err := u.db.WithContext(ctx).Create(&userRole).Error
	if err != nil {
		log.Printf("[UserRepo] Create userRole is failed %v: ", err)
		return err
	}
	return nil
}

func NewUserRoleRepository(db *gorm.DB) UserRoleRepository {
	return &userRoleRepository{
		db: db,
	}
}
