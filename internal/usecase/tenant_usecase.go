package usecase

import (
	"errors"
	"github.com/Prototype-1/Multi-Tenant-System/internal/model"
	"github.com/Prototype-1/Multi-Tenant-System/internal/repository"
)

type TenantUsecase interface {
	CreateTenant(name string) (*model.Tenant, error)
}

type tenantUsecase struct {
	tenantRepo repository.TenantRepository
}

func NewTenantUsecase(tenantRepo repository.TenantRepository) TenantUsecase {
	return &tenantUsecase{tenantRepo: tenantRepo}
}

func (u *tenantUsecase) CreateTenant(name string) (*model.Tenant, error) {
	existing, _ := u.tenantRepo.GetTenantByName(name)
	if existing != nil {
		return nil, errors.New("tenant with this name already exists")
	}

	tenant := &model.Tenant{
		Name: name,
	}

	err := u.tenantRepo.CreateTenant(tenant)
	if err != nil {
		return nil, err
	}
	return tenant, nil
}
