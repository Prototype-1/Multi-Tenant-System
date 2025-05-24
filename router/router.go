package router

import (
	"github.com/Prototype-1/Multi-Tenant-System/internal/handler"
	"github.com/Prototype-1/Multi-Tenant-System/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userHandler *handler.UserHandler,
	locationHandler *handler.LocationHandler,
	tenantHandler *handler.TenantHandler,
) *gin.Engine {
	r := gin.Default()
	r.POST("/create/tenants", tenantHandler.CreateTenant)

	r.POST("/users/signup", userHandler.Signup)
	r.POST("/users/login", userHandler.Login)

	userRoutes := r.Group("/get/users")
	userRoutes.Use(middleware.AuthMiddleware(), middleware.AuthorizeRole("admin"))
	userRoutes.GET("", userHandler.GetUsersHandler)

	r.POST("/create/locations", middleware.AuthMiddleware(), middleware.AuthorizeRole("user"), locationHandler.CreateLocation)
	r.GET("/get/me", middleware.AuthMiddleware(), middleware.AuthorizeRole("user"), userHandler.GetMe)

	return r
}
