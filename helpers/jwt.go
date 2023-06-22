package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(id, username, role, secret string) (string, error) {

	claims := jwt.MapClaims{
		"sub":      id,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

type PayloadJWT struct {
	ID       string
	Username string
	RoleID   string
}

func VerifyJWT(tokenString, secret string) (*PayloadJWT, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		fmt.Print(err)
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["sub"].(string)
	userName := claims["username"].(string)
	role := claims["role"].(string)

	data := PayloadJWT{

		ID:       userID,
		Username: userName,
		RoleID:   role,
	}

	return &data, nil
}
