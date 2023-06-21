package router

import (
	"trb-backend/module/admin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoutes(r *gin.Engine, db *gorm.DB) {
	adminHandler := admin.DefaultRequestAdminHandler(db)
	admins := r.Group("/admin")
	{
		admins.GET("/users", adminHandler.GetAllUser)
		admins.GET("/role/:id", adminHandler.GetAccessUser)
		admins.PUT("/role/:id", adminHandler.UpdateAccessUser)
	}
}
