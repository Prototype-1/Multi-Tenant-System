package pkg

import (
	"fmt"
	"log"
	"github.com/Prototype-1/Multi-Tenant-System/internal/model"
 "github.com/Prototype-1/Multi-Tenant-System/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := AutoMigrate(DB); err != nil {
		return nil, err
	}

	return DB, nil
}

func AutoMigrate(db *gorm.DB, models ...interface{}) error {
    if len(models) == 0 {
        models = []interface{}{
            &model.User{},
            &model.Location{},
			&model.Tenant{},
        }
        log.Println("Using default models for AutoMigrate")
    }
    
    err := db.AutoMigrate(models...)
    if err != nil {
        log.Printf("AutoMigrate error: %v", err)
        return fmt.Errorf("failed to migrate tables to database: %w", err)
    }
    log.Println("Database migration successful")
    return nil
}