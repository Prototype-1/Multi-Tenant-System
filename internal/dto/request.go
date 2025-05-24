package dto

import (
		"github.com/google/uuid"
		"time"
)

type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	TenantID string `json:"tenant_id" binding:"required,uuid"`
	Role     string `json:"role" binding:"required,oneof=admin user"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LocationResponse struct {
    ID        uuid.UUID `json:"id"`
    Latitude  float64   `json:"latitude"`
    Longitude float64   `json:"longitude"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type UserMeResponse struct {
    ID        uuid.UUID          `json:"id"`
    TenantID  uuid.UUID          `json:"tenant_id"`
    Email     string             `json:"email"`
    Role      string             `json:"role"`
    Locations []LocationResponse `json:"locations"`
    CreatedAt time.Time          `json:"created_at"`
    UpdatedAt time.Time          `json:"updated_at"`
}