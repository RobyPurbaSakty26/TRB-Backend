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
		admins.GET("/transactions", middleware.AccessMiddleware("Monitoring", "read", db),
			adminHandler.GetAllTransaction)
		admins.GET("/transactions-filter-by-date", middleware.AccessMiddleware("Download", "read", db),
			adminHandler.GetVritualAccountByDate)
		adminSecure := admins.Use(middleware.AdminAuthorization)
		{
			adminSecure.GET("/users", adminHandler.GetAllUsers)
			adminSecure.POST("/role", adminHandler.CreateRole)
			adminSecure.GET("/role/:roleId", adminHandler.GetListAccessRole)
			adminSecure.GET("/roles", adminHandler.GetAllRoles)
			adminSecure.PUT("/role/:roleId", adminHandler.UpdateAccessRole)
			adminSecure.DELETE("/role/:roleId", adminHandler.DeleteRole)
			adminSecure.PATCH("/active/:userId", adminHandler.UserApprove)
			adminSecure.DELETE("/user/:userId", adminHandler.DeleteUser)
			adminSecure.PUT("/user/role/:userId", adminHandler.AssignRole)
			adminSecure.GET("/accesses", adminHandler.GetListAccessName)
		}
	}
}
