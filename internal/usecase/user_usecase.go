package usecase

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"github.com/Prototype-1/Multi-Tenant-System/internal/dto"
	"github.com/Prototype-1/Multi-Tenant-System/internal/model"
	"github.com/Prototype-1/Multi-Tenant-System/internal/repository"
	"github.com/Prototype-1/Multi-Tenant-System/pkg"
)

type UserUsecase interface {
	Signup(req dto.SignupRequest) error
	Login(req dto.LoginRequest) (string, error)
	GetUsersByTenant(tenantID uuid.UUID) ([]model.User, error) 
	GetMeByID(userID uuid.UUID) (*model.User, error) 
}

type userUsecase struct {
	userRepo repository.UserRepository
	tenantRepo repository.TenantRepository
}

func NewUserUsecase(userRepo repository.UserRepository, tenantRepo repository.TenantRepository) UserUsecase {
	return &userUsecase{
		userRepo,
		tenantRepo,
	}
}

func (u *userUsecase) Signup(req dto.SignupRequest) error {

	if req.TenantID == "" {
        return errors.New("tenant ID is required")
    }

	tenantUUID, err := uuid.Parse(req.TenantID)
	if err != nil {
		return errors.New("invalid tenant ID")
	}

   tenant, err := u.tenantRepo.GetTenantByID(tenantUUID.String())
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return errors.New("tenant not found")
        }
        return errors.New("failed to validate tenant")
    }
    if tenant == nil {
        return errors.New("tenant not found")
    }

	existingUser, err := u.userRepo.FindByEmail(req.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("email is already registered")
	}
	if req.Role == "admin" {
		count, err := u.userRepo.CountAdminsByTenant(tenantUUID)
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("admin already exists for tenant")
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		TenantID: tenantUUID,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	return u.userRepo.Create(user)
}

func (u *userUsecase) Login(req dto.LoginRequest) (string, error) {
	user, err := u.userRepo.FindByEmail(req.Email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

  token, err := pkg.GenerateAccessToken(
        user.ID.String(),      
        user.TenantID.String(), 
        user.Role,
    )
    if err != nil {
        return "", err
    }
    return token, nil
}

func (u *userUsecase) GetUsersByTenant(tenantID uuid.UUID) ([]model.User, error) {
	return u.userRepo.FindUsersByTenant(tenantID)
}

func (u *userUsecase) GetMeByID(userID uuid.UUID) (*model.User, error) {
	return u.userRepo.GetMeById(userID)
}


