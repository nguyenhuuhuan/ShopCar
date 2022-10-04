package repositories

import (
	"Improve/src/models"
	"context"
	"errors"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(ctx context.Context, role *models.Role) error
	GetByID(ctx context.Context, id int64) (*models.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func (r roleRepository) Create(ctx context.Context, role *models.Role) error {
	err := r.db.WithContext(ctx).Create(&role).Error
	if errors.Is(err, gorm.ErrInvalidDB) {
		return err
	}
	return nil
}

func (r roleRepository) GetByID(ctx context.Context, id int64) (*models.Role, error) {
	var role *models.Role
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&role, 1).Error
	if errors.Is(err, gorm.ErrInvalidDB) {
		return nil, err
	}
	return role, nil
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}
