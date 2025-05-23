package usecase

import (
	"github.com/Prototype-1/Multi-Tenant-System/internal/model"
	"github.com/Prototype-1/Multi-Tenant-System/internal/repository"
)

type LocationUsecase interface {
	CreateLocation(location *model.Location) error
}

type locationUsecase struct {
	locationRepo repository.LocationRepository
}

func NewLocationUsecase(r repository.LocationRepository) LocationUsecase {
	return &locationUsecase{locationRepo: r}
}

func (u *locationUsecase) CreateLocation(location *model.Location) error {
	return u.locationRepo.Save(location)
}
