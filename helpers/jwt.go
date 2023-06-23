package helpers

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(id, username, roleId, roleName, secret string) (string, error) {

	claims := jwt.MapClaims{
		"sub":       id,
		"username":  username,
		"role_id":   roleId,
		"role_name": roleName,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
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
	RoleName string
}

func VerifyJWT(tokenString, secret string) (*PayloadJWT, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("terjadi panic")
		}
		return
	}()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	fmt.Println("data ++++++++++++ ", token.Valid)

	claims := token.Claims.(jwt.MapClaims)

	userID := claims["sub"].(string)
	userName := claims["username"].(string)
	roleId := claims["role_id"].(string)
	roleName := claims["role_name"].(string)

	data := PayloadJWT{
		ID:       userID,
		Username: userName,
		RoleID:   roleId,
		RoleName: roleName,
	}

	return &data, nil
}
