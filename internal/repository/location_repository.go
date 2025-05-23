package repository

import (
	"gorm.io/gorm"
	"github.com/Prototype-1/Multi-Tenant-System/internal/model"
)

type LocationRepository interface {
	Save(location *model.Location) error
}

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepository{db: db}
}

func (r *locationRepository) Save(location *model.Location) error {
	return r.db.Create(location).Error
}
