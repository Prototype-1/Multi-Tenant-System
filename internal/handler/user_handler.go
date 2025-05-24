package handler

import (
	"net/http"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/Multi-Tenant-System/internal/dto"
	"github.com/Prototype-1/Multi-Tenant-System/internal/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase}
}

func (h *UserHandler) Signup(c *gin.Context) {
	var req dto.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userUsecase.Signup(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful, please log in"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userUsecase.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
    "message": "Login successful",
    "token":   token,
    "status":  "success",
})
}

func (h *UserHandler) GetUsersHandler(c *gin.Context) {
	tenantIDStr, exists := c.Get("tenant_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a valid tenant ID"})
		return
	}

	tenantID, err := uuid.Parse(tenantIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	users, err := h.userUsecase.GetUsersByTenant(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetMe(c *gin.Context) {
    userID, ok := c.Get("user_id")
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID incorrect or missing"})
        return
    }
    
    userIDStr := userID.(string)  
    uid, err := uuid.Parse(userIDStr)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
        return
    }
    
    user, err := h.userUsecase.GetMeByID(uid)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user data"})
        return
    }
    
  var locationResponses []dto.LocationResponse
    for _, loc := range user.Locations {
        locationResponses = append(locationResponses, dto.LocationResponse{
            ID:        loc.ID,
            Latitude:  loc.Latitude,
            Longitude: loc.Longitude,
            CreatedAt: loc.CreatedAt,
            UpdatedAt: loc.UpdatedAt,
        })
    }
    
    response := dto.UserMeResponse{
        ID:        user.ID,
        TenantID:  user.TenantID,
        Email:     user.Email,
        Role:      user.Role,
        Locations: locationResponses,
        CreatedAt: user.CreatedAt,
        UpdatedAt: user.UpdatedAt,
    }
    
    c.JSON(http.StatusOK, response)
}

