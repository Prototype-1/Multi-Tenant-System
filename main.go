package main

import (
	"log"

	"github.com/Prototype-1/Multi-Tenant-System/config"
	"github.com/Prototype-1/Multi-Tenant-System/internal/handler"
	"github.com/Prototype-1/Multi-Tenant-System/internal/repository"
	"github.com/Prototype-1/Multi-Tenant-System/internal/usecase"
	"github.com/Prototype-1/Multi-Tenant-System/pkg"
	"github.com/Prototype-1/Multi-Tenant-System/router"
)

func main() {
	config.LoadConfig()

	db, err := pkg.InitDB(config.AppConfig)
	if err != nil {
		log.Fatalf("Failed to initialize DB: %v", err)
	}
	userRepo := repository.NewUserRepository(db)
	tenantRepo := repository.NewTenantRepository(db)
	locationRepo := repository.NewLocationRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo)
	tenantUsecase := usecase.NewTenantUsecase(tenantRepo)
	locationUsecase := usecase.NewLocationUsecase(locationRepo)

	userHandler := handler.NewUserHandler(userUsecase)
	tenantHandler := handler.NewTenantHandler(tenantUsecase)
	locationHandler := handler.NewLocationHandler(locationUsecase)

	r := router.SetupRouter(userHandler, locationHandler, tenantHandler)

	log.Printf("Server running on port %s...", config.AppConfig.Port)
	if err := r.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
