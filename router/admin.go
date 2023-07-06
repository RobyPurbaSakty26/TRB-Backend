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
	}
	roleReadAccess := r.Group("/admin").Use(middleware.AuthMiddleware).
		Use(middleware.AccessMiddleware("Role", "read", db))
	{
		roleReadAccess.GET("/role/:roleId", adminHandler.GetListAccessRole)
		roleReadAccess.GET("/roles", adminHandler.GetAllRoles)
		roleReadAccess.GET("/accesses", adminHandler.GetListAccessName)
	}
	roleWriteAccess := r.Group("/admin").Use(middleware.AuthMiddleware).
		Use(middleware.AccessMiddleware("Role", "write", db))
	{
		roleWriteAccess.POST("/role", adminHandler.CreateRole)
		roleWriteAccess.PUT("/role/:roleId", adminHandler.UpdateAccessRole)
		roleWriteAccess.DELETE("/role/:roleId", adminHandler.DeleteRole)
	}
	userAccess := r.Group("/admin").Use(middleware.AuthMiddleware)
	{
		userAccess.GET("/users", middleware.AccessMiddleware("User", "read", db),
			adminHandler.GetAllUsers)
		writeUser := userAccess.Use(middleware.AccessMiddleware("User", "write", db))
		{
			writeUser.PATCH("/active/:userId", adminHandler.UserApprove)
			writeUser.DELETE("/user/:userId", adminHandler.DeleteUser)
			writeUser.PUT("/user/role/:userId", adminHandler.AssignRole)
		}
	}
}
