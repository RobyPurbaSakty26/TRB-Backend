package middleware

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
	"strings"
	"trb-backend/helpers"
	"trb-backend/module/entity"
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

	if token == nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Token Not valid", Status: "False"})
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

	if data.RoleName != "admin" && data.RoleName != "Admin" {
		fmt.Println("jalan")
		c.JSON(http.StatusUnauthorized,
			response.ErrorResponse{
				Status:  "Fail",
				Message: "You are not allowed to access this page"})
		c.Abort()
		return
	}
	c.Next()
}

func AccessMiddleware(resource, permission string, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, _ := c.Get("data")
		data, _ := claim.(helpers.PayloadJWT)
		idUint64, _ := strconv.ParseUint(data.RoleID, 10, 64)
		idUint := uint(idUint64)
		var access entity.Access
		err := db.Where(entity.Access{Resource: resource, RoleId: idUint}).
			First(&access).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				response.ErrorResponse{
					Status:  "Failed",
					Message: err.Error()})
			return
		}

		if permission == "read" && access.CanRead != true {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				response.ErrorResponse{
					Status:  "Failed",
					Message: "You are not allowed to access this page"})
			return
		}
		if permission == "write" && access.CanWrite != true {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				response.ErrorResponse{
					Status:  "Failed",
					Message: "You are not allowed to access this page"})
			return
		}
		c.Next()
	}
}
