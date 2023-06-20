package helpers

import (
	"fmt"
	"time"

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
	Username string
	RoleID   string
}

func VerifyJWT(tokenString, secret string) (*PayloadJWT, error) {
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
		Username: userName,
		RoleID:   role,
	}

	return &data, nil
}
