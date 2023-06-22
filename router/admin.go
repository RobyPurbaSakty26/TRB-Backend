package router

import (
	"trb-backend/module/admin"
	"trb-backend/module/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoutes(r *gin.Engine, db *gorm.DB) {
	adminHandler := admin.DefaultRequestAdminHandler(db)
	admins := r.Group("/admin").Use(middleware.AuthMiddleware)
	{
		admins.GET("/users", middleware.AdminAuthorization, adminHandler.GetAllUsers)
		admins.GET("/role/:id", adminHandler.GetAccessUser)
		admins.PUT("/role/:id", adminHandler.UpdateAccessUser)
		admins.PATCH("/active/:id", adminHandler.UserApprove)
		admins.DELETE("/user/:id", adminHandler.DeleteUser)
	}
}
