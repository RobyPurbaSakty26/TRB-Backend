package admin

import (
	"trb-backend/module/entity"
	"trb-backend/module/web/request"
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

type useCase struct {
	repo AdminRepositoryInterface
}

type UseCaseAdminInterface interface {
	getAllUser(offset, limit int) ([]entity.User, error)
	getAllRoles(offset, limit int) ([]entity.Role, error)
	createRole(req *entity.Role) error
	updateAccess(req *entity.Access, id uint) error
	getAllAccessByRoleId(id string) ([]entity.Access, error)
	getRoleById(id string) (*entity.Role, error)
	updateRole(role *entity.Role, id uint) error
	userApprove(user *entity.User) error
	getById(id uint) (*entity.User, error)
	deleteUser(id uint) error
	createAccess(access *entity.Access) error
	deleteAccess(id uint) error
	deleteRole(id string) error
	assignRole(roleId uint, userId string) error
	getAllTransaction(offset, limit int) ([]entity.MasterAccount, error)
	getListAccess() ([]string, error)
	findVirtualAccountByDate(req *request.FillterTransactionByDate) ([]entity.TransactionVirtualAccount, error)
	findGiroByDate(req *request.FillterTransactionByDate) ([]entity.TransactionAccount, error)
	TotalDataMaster() (int64, error)
	TotalDataRole() (int64, error)
	TotalDataUser() (int64, error)
	findGiroByDatePagination(req *request.FillterTransactionByDate) ([]entity.TransactionAccount, error)
	findVaByDatePagination(req *request.FillterTransactionByDate) ([]entity.TransactionVirtualAccount, error)
	TotalDataTransactionGiro(req *request.FillterTransactionByDate) (int64, error)
	TotalDataTransactionVa(req *request.FillterTransactionByDate) (int64, error)
}

func NewUseCase(repo AdminRepositoryInterface) UseCaseAdminInterface {
	return useCase{
		repo: repo,
	}
}

func (u useCase) TotalDataTransactionGiro(req *request.FillterTransactionByDate) (int64, error) {
	return u.repo.TotalDataTransactionGiro(req)
}

func (u useCase) TotalDataTransactionVa(req *request.FillterTransactionByDate) (int64, error) {
	return u.repo.TotalDataTransactionGiro(req)
}

func (u useCase) TotalDataUser() (int64, error) {
	return u.repo.TotalDataUser()
}
func (u useCase) TotalDataRole() (int64, error) {
	return u.repo.TotalDataRole()
}
func (u useCase) TotalDataMaster() (int64, error) {
	return u.repo.TotalDataMaster()
}

func (u useCase) findGiroByDatePagination(req *request.FillterTransactionByDate) ([]entity.TransactionAccount, error) {
	return u.repo.getGiroByDatePagination(req)
}

func (u useCase) findVaByDatePagination(req *request.FillterTransactionByDate) ([]entity.TransactionVirtualAccount, error) {
	return u.repo.getVaByDatePagination(req)
}

func (u useCase) findGiroByDate(req *request.FillterTransactionByDate) ([]entity.TransactionAccount, error) {
	return u.repo.getGiroByDate(req)
}

func (u useCase) findVirtualAccountByDate(req *request.FillterTransactionByDate) ([]entity.TransactionVirtualAccount, error) {
	return u.repo.getVirtualAccountByDate(req)
}

func (u useCase) getListAccess() ([]string, error) {
	return u.repo.getListAccess()
}
func (u useCase) getAllTransaction(offset, limit int) ([]entity.MasterAccount, error) {
	return u.repo.getAllTransaction(offset, limit)
}
func (u useCase) assignRole(roleId uint, userId string) error {
	return u.repo.assignRole(roleId, userId)
}
func (u useCase) getAllRoles(offset, limit int) ([]entity.Role, error) {
	return u.repo.getAllRoles(offset, limit)
}
func (u useCase) createRole(req *entity.Role) error {
	return u.repo.createRole(req)
}

func (u useCase) createAccess(access *entity.Access) error {
	return u.repo.createAccess(access)
}

func (u useCase) getAllUser(offset, limit int) ([]entity.User, error) {
	return u.repo.getAllUser(offset, limit)
}

func (u useCase) updateAccess(req *entity.Access, id uint) error {
	return u.repo.updateAccess(req, id)
}

func (u useCase) getRoleById(id string) (*entity.Role, error) {
	return u.repo.getRoleById(id)
}

func (u useCase) getAllAccessByRoleId(id string) ([]entity.Access, error) {
	return u.repo.getAllAccessByRoleId(id)
}

func (u useCase) deleteAccess(id uint) error {
	return u.repo.deleteAccess(id)
}

func (u useCase) deleteRole(id string) error {
	return u.repo.deleteRole(id)
}

func (u useCase) updateRole(role *entity.Role, id uint) error {
	return u.repo.updateRole(role, id)
}
func (u useCase) userApprove(user *entity.User) error {
	return u.repo.userApprove(user)
}

func (u useCase) getById(id uint) (*entity.User, error) {
	return u.repo.getById(id)
}

func (u useCase) deleteUser(id uint) error {
	// Periksa apakah pengguna dengan ID tersebut ada dalam sistem
	_, err := u.repo.getById(id)
	if err != nil {
		return err
	}

	// Melakukan penghapusan pengguna dari repository
	err = u.repo.deleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
