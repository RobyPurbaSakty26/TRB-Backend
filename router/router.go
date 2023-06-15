package router

import (
	"log"
	"trb-backend/config"
	"trb-backend/module/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()

	userHandler := user.DefaultRequestHandler(db)

	r.POST("/register", userHandler.Create)
	r.GET("/user/email", userHandler.GetByEmail)
	r.GET("/user/username", userHandler.GetByUsername)
	r.POST("/login", userHandler.Login)
	r.PATCH("/user/forgot-password", userHandler.UpdatePassword)

	return r
}
