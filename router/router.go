package router

import (
	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Authorization"},
		AllowCredentials: true,
	}))

	userHandler := user.DefaultRequestHandler(db)

	r.POST("/register", userHandler.Create)
	r.GET("/user/email", userHandler.GetByEmail)
	r.GET("/user/username", userHandler.GetByUsername)
	r.POST("/login", userHandler.Login)
	r.PATCH("/user/forgot-password", userHandler.UpdatePassword)

	return r
}
