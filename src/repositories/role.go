package repositories

import (
	"Improve/src/models"
	"context"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(ctx context.Context, role *models.Role) error
	GetByID(ctx context.Context, id int64) (*models.Role, error)
	IsExistRoleAndCode(ctx context.Context, roleName, code string) bool
}

type roleRepository struct {
	db *gorm.DB
}

func (r roleRepository) IsExistRoleAndCode(ctx context.Context, roleName, code string) bool {
	var role *models.Role
	if err := r.db.WithContext(ctx).Where("role_name = ? or code = ?", roleName, code).Take(&role).Error; err != nil {
		return false
	}
	return true
}

func (r roleRepository) Create(ctx context.Context, role *models.Role) error {
	err := r.db.WithContext(ctx).Create(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func (r roleRepository) GetByID(ctx context.Context, id int64) (*models.Role, error) {
	var role *models.Role
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&role, 1).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}
