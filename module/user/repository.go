package user

import (
	"trb-backend/module/entity"

	"gorm.io/gorm"
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

type repository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	save(user *entity.User) error
	getByEmail(email string) (*entity.User, error)
	getByUsername(username string) (*entity.User, error)
	getUserAndRole(id uint) (*entity.User, error)
	updatePassword(user *entity.User, password string) error
	updateInputFalse(user *entity.User, count int) error
	updateStatusIsActive(user *entity.User, isActive bool) error
	getById(id int) (*entity.User, error)
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
	err := r.db.Where("username = ?", username).Preload("Role").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r repository) getUserAndRole(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r repository) updatePassword(user *entity.User, password string) error {
	return r.db.Model(user).Where("email = ?", user.Email).Update("password", password).Error
}

func (r repository) updateInputFalse(user *entity.User, count int) error {
	return r.db.Model(user).Where(" email = ? ", user.Email).Update("input_false", count).Error
}

func (r repository) updateStatusIsActive(user *entity.User, isActive bool) error {
	return r.db.Model(user).Where("email = ?", user.Email).Update("active", isActive).Error
}

func (r repository) getById(id int) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error

	return &user, err

}
