package router

import (
	"log"
	"trb-backend/config"
	"trb-backend/module/middleware"
	"trb-backend/module/user"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

/**
 * Created by Goland & VS Code.
 * User : 1. Roby Purba Sakty 			: obykao26@gmail.com
		  2. Muhammad Irfan 			: mhd.irfann00@gmail.com
   		  3. Andre Rizaldi Brillianto	: andrerizaldib@gmail.com
 * Date: Saturday, 12 Juni 2023
 * Time: 08.30 AM
 * Description: BRI-CMP-Service-Backend
 **/

func SetupRouter() *gin.Engine {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	}))

	userHandler := user.DefaultRequestHandler(db)

	AdminRoutes(r, db)
	r.POST("/register", userHandler.Create)
	r.POST("/login", userHandler.Login)
	r.PATCH("/user/forgot-password", userHandler.UpdatePassword)
	// auth
	r.GET("/user", middleware.AuthMiddleware, userHandler.GetAllUsers)
	r.GET("/user/email", middleware.AuthMiddleware, userHandler.GetByEmail)
	r.GET("/user/username", middleware.AuthMiddleware, userHandler.GetByUsername)
	r.PATCH("/user/approve/:id", middleware.AuthMiddleware, userHandler.UserApprove)

	return r
}
