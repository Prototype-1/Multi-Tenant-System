package model

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tenant struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null" json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Tenant) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
