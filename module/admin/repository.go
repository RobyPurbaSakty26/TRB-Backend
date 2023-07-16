package admin

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

type AdminRepositoryInterface interface {
	getAllUser(offset, limit int) ([]entity.User, error)
	getAllRoles(offset, limit int) ([]entity.Role, error)
	getAllAccessByRoleId(id string) ([]entity.Access, error)
	getRoleById(id string) (*entity.Role, error)
	updateAccess(request *entity.Access, id uint) error
	updateRole(role *entity.Role, id uint) error
	userApprove(user *entity.User) error
	getById(id uint) (*entity.User, error)
	deleteUser(id uint) error
	createRole(req *entity.Role) error
	createAccess(access *entity.Access) error
	deleteRole(id string) error
	deleteAccess(id uint) error
	assignRole(roleId uint, userId string) error
	getAllTransaction(offset, limit int) ([]entity.MasterAccount, error)
	getListAccess() ([]string, error)
	getVirtualAccountByDate(accNo, startDate, endDate string) ([]entity.TransactionVirtualAccount, error)
	getGiroByDate(accNo, startDate, endDate string) ([]entity.TransactionAccount, error)
	TotalDataMaster() (int64, error)
	TotalDataRole() (int64, error)
	TotalDataUser() (int64, error)
	getGiroByDatePagination(accNo, startDate, endDate string, limit, page int) ([]entity.TransactionAccount, error)
	getVaByDatePagination(accNo, startDate, endDate string, limit, page int) ([]entity.TransactionVirtualAccount, error)
	TotalDataTransactionGiro(accNo, startDate, endDate string) (int64, error)
	TotalDataTransactionVa(accNo, startDate, endDate string) (int64, error)
	getUserByUsername(username string, page, limit int) ([]entity.User, error)
	totalGetUserByUsername(username string) (int64, error)
	totalGetUserByEmail(email string) (int64, error)
	getUserByEmail(email string, page, limit int) ([]entity.User, error)
}

func NewAdminRepository(db *gorm.DB) AdminRepositoryInterface {
	return &repository{db: db}
}

func (r repository) totalGetUserByUsername(username string) (int64, error) {
	var count int64
	err := r.db.Table("users").Where("username LIKE ? AND deleted_at IS NULL", "%"+username+"%").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r repository) totalGetUserByEmail(email string) (int64, error) {
	var count int64
	err := r.db.Table("users").Where("email LIKE ? AND deleted_at IS NULL", "%"+email+"%").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r repository) getUserByUsername(username string, page, limit int) ([]entity.User, error) {
	var users []entity.User
	err := r.db.Where("username LIKE ?", "%"+username+"%").Limit(limit).Offset(page).Preload("Role").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r repository) getUserByEmail(email string, page, limit int) ([]entity.User, error) {
	var users []entity.User
	err := r.db.Where("email LIKE ? ", "%"+email+"%").Limit(limit).Offset(page).Preload("Role").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r repository) TotalDataUser() (int64, error) {
	var count int64
	err := r.db.Table("users").Where("deleted_at IS NULL").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r repository) TotalDataRole() (int64, error) {
	var count int64
	err := r.db.Table("roles").Where("deleted_at is NULL").Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r repository) TotalDataMaster() (int64, error) {
	var count int64
	err := r.db.Table("master_account").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r repository) TotalDataTransactionGiro(accNo, startDate, endDate string) (int64, error) {
	var count int64
	err := r.db.Table("transaction_account").Where("account_no = ? AND (transaction_date >= ? AND transaction_date <= ?)", accNo, startDate, endDate).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r repository) TotalDataTransactionVa(accNo, startDate, endDate string) (int64, error) {
	var count int64
	err := r.db.Table("transaction_virtual_account").Where("account_no = ? AND (transaction_date >= ? AND transaction_date <= ?)", accNo, startDate, endDate).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r repository) getGiroByDatePagination(accNo, startDate, endDate string, limit, page int) ([]entity.TransactionAccount, error) {
	var datas []entity.TransactionAccount
	err := r.db.Where("account_no = ? AND (transaction_date >= ? AND transaction_date <= ?)", accNo, startDate, endDate).Limit(limit).Offset(page).Find(&datas).Error
	if err != nil {
		return nil, err
	}
	return datas, err
}

func (r repository) getVaByDatePagination(accNo, startDate, endDate string, limit, page int) ([]entity.TransactionVirtualAccount, error) {
	var datas []entity.TransactionVirtualAccount
	err := r.db.Where("account_no = ? AND (transaction_date >= ? AND transaction_date <= ?)", accNo, startDate, endDate).Limit(limit).Offset(page).Find(&datas).Error
	if err != nil {
		return nil, err
	}
	return datas, err
}

func (r repository) getGiroByDate(accNo, startDate, endDate string) ([]entity.TransactionAccount, error) {
	var datas []entity.TransactionAccount
	err := r.db.Where("account_no = ? AND (transaction_date >= ? AND transaction_date <= ?)", accNo, startDate, endDate).Find(&datas).Error
	if err != nil {
		return nil, err
	}
	return datas, err
}

func (r repository) getVirtualAccountByDate(accNo, startDate, endDate string) ([]entity.TransactionVirtualAccount, error) {
	var datas []entity.TransactionVirtualAccount
	err := r.db.Where("account_no = ? AND (transaction_date >= ? AND transaction_date <= ?)", accNo, startDate, endDate).Find(&datas).Error
	if err != nil {
		return nil, err
	}
	return datas, nil
}

func (r repository) getListAccess() ([]string, error) {
	var names []string
	var access entity.Access
	err := r.db.Model(&access).Select("DISTINCT resource").Find(&names).Error
	if err != nil {
		return nil, err
	}
	return names, nil
}
func (r repository) getAllTransaction(offset, limit int) ([]entity.MasterAccount, error) {
	var datas []entity.MasterAccount
	err := r.db.Limit(limit).Offset(offset).Find(&datas).Error

	if err != nil {
		return nil, err
	}
	return datas, nil
}

func (r repository) assignRole(roleId uint, userId string) error {
	var user entity.User
	return r.db.Model(&user).
		Where("id = ?", userId).
		Update("role_id", roleId).Error
}
func (r repository) getAllRoles(offset, limit int) ([]entity.Role, error) {
	var roles []entity.Role
	err := r.db.Limit(limit).Offset(offset).Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}
func (r repository) createRole(req *entity.Role) error {
	return r.db.Create(req).Error
}

func (r repository) createAccess(access *entity.Access) error {
	return r.db.Create(access).Error
}

func (r repository) getAllUser(offset, limit int) ([]entity.User, error) {
	var user []entity.User
	err := r.db.Offset(offset).Limit(limit).Preload("Role").Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r repository) getRoleById(id string) (*entity.Role, error) {
	var role entity.Role
	err := r.db.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}
func (r repository) getAllAccessByRoleId(id string) ([]entity.Access, error) {
	var access []entity.Access
	err := r.db.Find(&access, "role_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return access, nil
}
func (r repository) updateRole(role *entity.Role, id uint) error {
	return r.db.Model(role).
		Where("id = ?", id).
		Update("name", role.Name).Error
}
func (r repository) updateAccess(request *entity.Access, id uint) error {
	return r.db.Model(request).
		Where(entity.Access{Resource: request.Resource, RoleId: id}).
		Updates(map[string]interface{}{
			"can_read":  request.CanRead,
			"can_write": request.CanWrite,
		}).Error
}

func (r repository) deleteAccess(id uint) error {
	var access entity.Access
	return r.db.Delete(&access, "role_id = ?", id).Error
}

func (r repository) deleteRole(id string) error {
	var role entity.Role
	return r.db.Delete(&role, id).Error
}
func (r repository) userApprove(user *entity.User) error {

	return r.db.Model(&user).Updates(map[string]interface{}{
		"InputFalse": 0,
		"Active":     true,
	}).Error
}

func (r repository) getById(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error

	return &user, err

}

func (r *repository) deleteUser(id uint) error {
	var user entity.User
	return r.db.Delete(&user, id).Error
}
