package repository

import (
	"github.com/Prototype-1/Multi-Tenant-System/internal/model"
	"gorm.io/gorm"
)

type TenantRepository interface {
	CreateTenant(tenant *model.Tenant) error
	GetTenantByName(name string) (*model.Tenant, error)
	GetTenantByID(id string) (*model.Tenant, error)
}

type tenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) TenantRepository {
	return &tenantRepository{db: db}
}

func (r *tenantRepository) CreateTenant(tenant *model.Tenant) error {
	return r.db.Create(tenant).Error
}

func (r *tenantRepository) GetTenantByName(name string) (*model.Tenant, error) {
	var tenant model.Tenant
	err := r.db.Where("name = ?", name).First(&tenant).Error
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *tenantRepository) GetTenantByID(id string) (*model.Tenant, error) {
	var tenant model.Tenant
	err := r.db.Where("id = ?", id).First(&tenant).Error
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}
