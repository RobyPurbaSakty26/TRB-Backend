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
		admins.GET("/transactions/download", middleware.AccessMiddleware("Monitoring", "write", db),
			adminHandler.DownloadTransaction)
		admins.GET("/transactions-filter-by-date", middleware.AccessMiddleware("Download", "read", db),
			adminHandler.GetTransactionByDate)
    admins.GET("/transactions-filter-by-date/download", middleware.AccessMiddleware("Download", "write", db),
			adminHandler.DownloadTransactionByDate)
		admins.GET("/users", middleware.AccessMiddleware("User", "read", db),
			adminHandler.GetAllUsers)
		writeUser := admins.Use(middleware.AccessMiddleware("User", "write", db))
		{
			writeUser.PATCH("/active/:userId", adminHandler.UserApprove)
			writeUser.DELETE("/user/:userId", adminHandler.DeleteUser)
			writeUser.PUT("/user/role/:userId", adminHandler.AssignRole)
		}
		readRole := admins.Use(middleware.AccessMiddleware("Role", "read", db))
		{
			readRole.GET("/role/:roleId", adminHandler.GetListAccessRole)
			readRole.GET("/roles", adminHandler.GetAllRoles)
			readRole.GET("/accesses", adminHandler.GetListAccessName)
		}
		writeRole := admins.Use(middleware.AccessMiddleware("Role", "write", db))
		{
			writeRole.POST("/role", adminHandler.CreateRole)
			writeRole.PUT("/role/:roleId", adminHandler.UpdateAccessRole)
			writeRole.DELETE("/role/:roleId", adminHandler.DeleteRole)
		}
	}
}
