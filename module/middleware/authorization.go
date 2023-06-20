package middleware

import (
	"net/http"
	"strings"
	"trb-backend/helpers"
	"trb-backend/module/web"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	// get token from authorization
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// verify token
	token, err := helpers.VerifyJWT(tokenString, "secret-key")
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, web.ErrorResponse{Status: "Fail", Message: err.Error()})
		c.Abort()
		return
	}

	data := helpers.PayloadJWT{
		ID:       token.ID,
		Username: token.Username,
		RoleID:   token.RoleID,
	}

	c.Set("data", data)

	// c.Set("id", token.ID)
	// c.Set("username", token.username)
	// c.Set("role", token.Role)
	c.Next()
}
