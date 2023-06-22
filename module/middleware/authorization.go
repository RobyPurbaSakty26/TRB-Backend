package middleware

import (
	"net/http"
	"os"
	"strings"
	"trb-backend/helpers"
	"trb-backend/module/web/response"

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

// generate token

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	secret := os.Getenv("SECRET_KEY")
	token, err := helpers.VerifyJWT(tokenString, secret)

	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, response.ErrorResponse{Status: "Fail", Message: err.Error()})
		c.Abort()
		return
	}

	data := helpers.PayloadJWT{
		ID:       token.ID,
		Username: token.Username,
		RoleID:   token.RoleID,
		RoleName: token.RoleName,
	}

	c.Set("data", data)

	c.Next()
}

func AdminAuthorization(c *gin.Context) {
	claim, _ := c.Get("data")
	data, _ := claim.(helpers.PayloadJWT)

	if data.RoleName != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not allowed to access this page",
		})
		return
	}
	c.Next()
}
