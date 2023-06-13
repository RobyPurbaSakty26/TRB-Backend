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

	return r
}
