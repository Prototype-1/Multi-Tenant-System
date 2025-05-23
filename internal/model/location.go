package model

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;index" json:"user_id"`
	TenantID  uuid.UUID `gorm:"type:uuid;index" json:"tenant_id"`
	Latitude  float64   `gorm:"not null" json:"latitude"`
	Longitude float64   `gorm:"not null" json:"longitude"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (l *Location) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = uuid.New()
	return
}
