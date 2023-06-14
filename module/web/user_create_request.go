package web

import (
	"errors"
	"gorm.io/gorm"
)

type UserCreateRequest struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserCreateRequest) BeforeCreate(tx *gorm.DB) (err error) {
	var user UserCreateRequest
	errInput := tx.Model(&user).Where("username = ? ", u.Username).First(&user).Error
	if errInput == nil {
		err = errors.New("username already exist")
	}
	errInput = tx.Model(&user).Where("email = ? ", u.Email).First(&user).Error
	if errInput == nil {
		err = errors.New("email already exist")
	}
	return
}
