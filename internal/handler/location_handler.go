package handler

import (
	"os"
	"fmt"
	"time"
	"net/http"
	"strings"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/Multi-Tenant-System/internal/model"
	"github.com/Prototype-1/Multi-Tenant-System/internal/usecase"
)

type LocationHandler struct {
	locationUsecase usecase.LocationUsecase
}

func NewLocationHandler(usecase usecase.LocationUsecase) *LocationHandler {
	return &LocationHandler{
		locationUsecase: usecase,
	}
}

type CreateLocationRequest struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

func (h *LocationHandler) CreateLocation(c *gin.Context) {
	var req CreateLocationRequest
   if err := c.ShouldBindJSON(&req); err != nil {
        if strings.Contains(err.Error(), "Latitude") || strings.Contains(err.Error(), "Longitude") {
            c.JSON(http.StatusBadRequest, gin.H{"error": "You need to provide both longitude and latitude to continue with this operation"})
            return
        }
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	role := c.MustGet("role").(string)
	if role == "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admins are not allowed to submit location data"})
		return
	}

	 userIDStr := c.MustGet("user_id").(string)
    tenantIDStr := c.MustGet("tenant_id").(string)
    
    userID, err := uuid.Parse(userIDStr)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
        return
    }
    
    tenantID, err := uuid.Parse(tenantIDStr)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid tenant ID"})
        return
    }

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

f, err := os.OpenFile("location_stream.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err == nil {
	defer f.Close()
	logLine := fmt.Sprintf("User %s submitted location (Lat: %.6f, Lon: %.6f) at %s\n",
		location.UserID, location.Latitude, location.Longitude, location.CreatedAt.Format(time.RFC3339))
	f.WriteString(logLine)
}

	   c.JSON(http.StatusCreated, gin.H{
        "message": "Location saved successfully",
        "location": location,
    })
}
