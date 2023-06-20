package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"trb-backend/module/admin"
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
