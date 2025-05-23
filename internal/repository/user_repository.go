package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/Prototype-1/Multi-Tenant-System/internal/model"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	CountAdminsByTenant(tenantID uuid.UUID) (int64, error)
	FindUsersByTenant(tenantID uuid.UUID) ([]model.User, error) 
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *userRepository) CountAdminsByTenant(tenantID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&model.User{}).Where("tenant_id = ? AND role = ?", tenantID, "admin").Count(&count).Error
	return count, err
}

func (r *userRepository) FindUsersByTenant(tenantID uuid.UUID) ([]model.User, error) {
	var users []model.User
	if err := r.db.Preload("Locations").Where("tenant_id = ?", tenantID).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

