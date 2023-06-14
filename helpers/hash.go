package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// start feature : register
func HashPass(p string) (string, error) {
	pass := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)

	return string(hash), err
}

func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}

//end feature : register
