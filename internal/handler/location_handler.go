package handler

import (
	"net/http"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/Multi-Tenant-System/internal/model"
	"github.com/Prototype-1/Multi-Tenant-System/internal/usecase"
)

type LocationHandler struct {
	locationUsecase usecase.LocationUsecase
}

type CreateLocationRequest struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

func (h *LocationHandler) CreateLocation(c *gin.Context) {
	var req CreateLocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := c.MustGet("role").(string)
	if role == "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admins are not allowed to submit location data"})
		return
	}

	userID := c.MustGet("user_id").(uuid.UUID)
	tenantID := c.MustGet("tenant_id").(uuid.UUID)

	location := model.Location{
		UserID:    userID,
		TenantID:  tenantID,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	if err := h.locationUsecase.CreateLocation(&location); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save location"})
		return
	}

	c.JSON(http.StatusCreated, location)
}
