package router

import (
	"trb-backend/module/admin"
	"trb-backend/module/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoutes(r *gin.Engine, db *gorm.DB) {
	adminHandler := admin.DefaultRequestAdminHandler(db)
	admins := r.Group("/admin").
		Use(middleware.AuthMiddleware).Use(middleware.AdminAuthorization)
	{
		admins.GET("/users", adminHandler.GetAllUsers)
		admins.POST("/role", adminHandler.CreateRole)
		admins.GET("/role/:id", adminHandler.GetListAccessRole)
		admins.GET("/roles", adminHandler.GetAllRoles)
		admins.PUT("/role/:id", adminHandler.UpdateAccessRole)
		admins.DELETE("/role/:id", adminHandler.DeleteRole)
		admins.PATCH("/active/:id", adminHandler.UserApprove)
		admins.DELETE("/user/:id", adminHandler.DeleteUser)
		admins.PUT("/user/role/:id", adminHandler.AssignRole)
	}
}
