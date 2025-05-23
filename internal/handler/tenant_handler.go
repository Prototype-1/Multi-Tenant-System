package handler

import (
	"net/http"
	"github.com/Prototype-1/Multi-Tenant-System/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TenantHandler struct {
	tenantUsecase usecase.TenantUsecase
}

func NewTenantHandler(tu usecase.TenantUsecase) *TenantHandler {
	return &TenantHandler{tenantUsecase: tu}
}

func (h *TenantHandler) CreateTenant(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required,min=3"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenant, err := h.tenantUsecase.CreateTenant(input.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"tenant": tenant})
}
