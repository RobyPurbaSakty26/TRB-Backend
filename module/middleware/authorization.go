package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"trb-backend/module/web"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// generate token
func GenerateToken(id, username, role, secret string) (string, error) {
	// ini sialisasi klaim
	claims := jwt.MapClaims{
		"sub":      id,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	// tandatangan token dengan kunci rahasia

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

type PayloadJWT struct {
	ID       string
	username string
	Role     string
}

func verifyJWT(tokenString, secret string) (*PayloadJWT, error) {
	// Memeriksa keaslian token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		fmt.Print(err)
		return nil, err
	}

	// Token valid, dapatkan informasi pengguna dari token
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["sub"].(string)
	userName := claims["username"].(string)
	role := claims["role"].(string)

	data := PayloadJWT{

		ID:       userID,
		username: userName,
		Role:     role,
	}

	return &data, nil
}

func AuthMiddleware(c *gin.Context) {
	// get token from authorization
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// verify token
	token, err := verifyJWT(tokenString, "secret-key")
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, web.ErrorResponse{Status: "Fail", Message: err.Error()})
		c.Abort()
		return
	}

	data := PayloadJWT{
		ID:       token.ID,
		username: token.username,
		Role:     token.Role,
	}

	c.Set("data", data)

	// c.Set("id", token.ID)
	// c.Set("username", token.username)
	// c.Set("role", token.Role)
	c.Next()
}
