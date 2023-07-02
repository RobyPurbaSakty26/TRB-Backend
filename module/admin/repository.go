package admin

import (
	"trb-backend/module/entity"
	"trb-backend/module/web/request"

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
	getAllUser() ([]entity.User, error)
	getAllRoles() ([]entity.Role, error)
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
	getAllTransaction(page, limit string) ([]entity.MasterAccount, error)
	getSaldoTransactionGiro(accNo string) (int, error)
	getSaldoTransactionVA(accNo string) (int, error)
	getTotalAccVA(accNo string) (int64, error)
	getListAccess() ([]string, error)
	getVirtualAccountByDate(req *request.FillterTransactionByDate) ([]entity.TransactionVirtualAccount, error)
	getGiroByDate(req *request.FillterTransactionByDate) ([]entity.TransactionAccount, error)
}

func NewAdminRepository(db *gorm.DB) AdminRepositoryInterface {
	return &repository{db: db}
}

func (r repository) getGiroByDate(req *request.FillterTransactionByDate) ([]entity.TransactionAccount, error) {
	var datas []entity.TransactionAccount
	err := r.db.Where("account_no = ? AND (transaction_date >= ? AND transaction_date <= ?)", req.AccNo, req.StartDate, req.EndDate).Find(&datas).Error
	if err != nil {
		return nil, err
	}
	return datas, err
}

func (r repository) getVirtualAccountByDate(req *request.FillterTransactionByDate) ([]entity.TransactionVirtualAccount, error) {
	var datas []entity.TransactionVirtualAccount
	err := r.db.Where("account_no = ? AND (transaction_date >= ? AND transaction_date <= ?)", req.AccNo, req.StartDate, req.EndDate).Find(&datas).Error
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
func (r repository) getTotalAccVA(accNo string) (int64, error) {
	var transaction entity.TransactionVirtualAccount
	var totalAcc int64
	err := r.db.Model(&transaction).Where("account_no = ?", accNo).Count(&totalAcc).Error
	if err != nil {
		return 0, err
	}
	return totalAcc, nil
}

func (r repository) getSaldoTransactionGiro(accNo string) (int, error) {
	var transaction entity.TransactionAccount
	var total, totalDebit, totalCredit int

	err := r.db.Model(&transaction).Select("Sum(amount)").
		Where(entity.TransactionAccount{Category: "Debit", AccountNo: accNo}).Scan(&totalDebit).Error
	if err != nil {
		return 0, err
	}
	err = r.db.Model(&transaction).Select("Sum(amount)").
		Where(entity.TransactionAccount{Category: "Credit", AccountNo: accNo}).Scan(&totalCredit).Error
	if err != nil {
		return 0, err
	}
	total = totalDebit - totalCredit
	return total, nil
}

func (r repository) getSaldoTransactionVA(accNo string) (int, error) {
	var transaction entity.TransactionVirtualAccount
	var total, totalDebit, totalCredit int

	err := r.db.Model(&transaction).Select("Sum(credit)").
		Where(entity.TransactionVirtualAccount{Category: "Debit", AccountNo: accNo}).Scan(&totalDebit).Error
	if err != nil {
		return 0, err
	}
	err = r.db.Model(&transaction).Select("Sum(credit)").
		Where(entity.TransactionVirtualAccount{Category: "Credit", AccountNo: accNo}).Scan(&totalCredit).Error
	if err != nil {
		return 0, err
	}
	total = totalDebit - totalCredit
	return total, nil
}

func (r repository) getAllTransaction(page, limit string) ([]entity.MasterAccount, error) {
	var datas []entity.MasterAccount
	err := r.db.Find(&datas).Error

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
func (r repository) getAllRoles() ([]entity.Role, error) {
	var roles []entity.Role
	err := r.db.Find(&roles).Error
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

func (r repository) getAllUser() ([]entity.User, error) {
	var user []entity.User
	err := r.db.Preload("Role").Find(&user).Error
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
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
