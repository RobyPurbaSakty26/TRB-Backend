package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"trb-backend/module/admin"
	"trb-backend/module/middleware"
)

func AdminRoutes(r *gin.Engine, db *gorm.DB) {
	adminHandler := admin.DefaultRequestAdminHandler(db)
	admins := r.Group("/admin").
		Use(middleware.AuthMiddleware).Use(middleware.AdminAuthorization)
	{
		admins.GET("/users", adminHandler.GetAllUsers)
		admins.POST("/role", adminHandler.CreateRole)
		admins.GET("/role/:roleId", adminHandler.GetListAccessRole)
		admins.GET("/roles", adminHandler.GetAllRoles)
		admins.PUT("/role/:roleId", adminHandler.UpdateAccessRole)
		admins.DELETE("/role/:roleId", adminHandler.DeleteRole)
		admins.PATCH("/active/:userId", adminHandler.UserApprove)
		admins.DELETE("/user/:userId", adminHandler.DeleteUser)
		admins.PUT("/user/role/:userId", adminHandler.AssignRole)
	}
	r.GET("/admin/transactions", adminHandler.GetAllTransaction)
}
