package admin

import (
	"gorm.io/gorm"
	"trb-backend/module/entity"
	"trb-backend/module/web"
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

type AdminRepositoryInterface interface {
	getAllUser() ([]entity.User, error)
	getAllAccessByRoleId(id string) ([]entity.Access, error)
	getUserWithRole(id string) (*entity.User, error)
	updateAccess(access *entity.Access, request *web.AccessRequest, id uint) error
}

func NewAdminRepository(db *gorm.DB) AdminRepositoryInterface {
	return &repository{db: db}
}

func (r repository) getAllUser() ([]entity.User, error) {
	var user []entity.User
	err := r.db.Preload("Role").Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r repository) getUserWithRole(id string) (*entity.User, error) {
	var users entity.User
	err := r.db.Preload("Role").First(&users, "role_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}
func (r repository) getAllAccessByRoleId(id string) ([]entity.Access, error) {
	var access []entity.Access
	err := r.db.Find(&access, "role_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return access, nil
}
func (r repository) updateAccess(access *entity.Access, request *web.AccessRequest, id uint) error {
	return r.db.Model(&access).
		Where(entity.Access{Resource: request.Resource, RoleId: id}).
		Updates(entity.Access{CanRead: request.CanRead, CanWrite: request.CanWrite}).Error
}
