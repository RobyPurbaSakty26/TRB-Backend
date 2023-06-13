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
}

func NewRepository(db *gorm.DB) UserRepositoryInterface {
	return &repository{db: db}
}

func (r repository) save(user *entity.User) error {
	return r.db.Create(user).Error
}
