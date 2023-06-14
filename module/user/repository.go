package user

import (
	"trb-backend/module/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	save(user *entity.User) error
	getByEmail(email string) (*entity.User, error)
	getByUsername(username string) (*entity.User, error)
}

func NewRepository(db *gorm.DB) UserRepositoryInterface {
	return &repository{db: db}
}

func (r repository) save(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r repository) getByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r repository) getByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
