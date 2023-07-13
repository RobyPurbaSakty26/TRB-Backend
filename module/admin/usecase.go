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
	GetAllUser(offset, limit int) ([]entity.User, error)
	GetAllRoles(offset, limit int) ([]entity.Role, error)
	CreateRole(req *entity.Role) error
	UpdateAccess(req *entity.Access, id uint) error
	GetAllAccessByRoleId(id string) ([]entity.Access, error)
	GetRoleById(id string) (*entity.Role, error)
	UpdateRole(role *entity.Role, id uint) error
	UserApprove(user *entity.User) error
	GetById(id uint) (*entity.User, error)
	DeleteUser(id uint) error
	CreateAccess(access *entity.Access) error
	DeleteAccess(id uint) error
	DeleteRole(id string) error
	AssignRole(roleId uint, userId string) error
	GetAllTransaction(offset, limit int) ([]entity.MasterAccount, error)
	GetListAccess() ([]string, error)
	FindVirtualAccountByDate(req *request.FillterTransactionByDate) ([]entity.TransactionVirtualAccount, error)
	FindGiroByDate(req *request.FillterTransactionByDate) ([]entity.TransactionAccount, error)
	TotalDataMaster() (int64, error)
	TotalDataRole() (int64, error)
	TotalDataUser() (int64, error)
	FindGiroByDatePagination(req *request.FillterTransactionByDate) ([]entity.TransactionAccount, error)
	FindVaByDatePagination(req *request.FillterTransactionByDate) ([]entity.TransactionVirtualAccount, error)
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

func (u useCase) FindGiroByDatePagination(req *request.FillterTransactionByDate) ([]entity.TransactionAccount, error) {
	return u.repo.GetGiroByDatePagination(req)
}

func (u useCase) FindVaByDatePagination(req *request.FillterTransactionByDate) ([]entity.TransactionVirtualAccount, error) {
	return u.repo.GetVaByDatePagination(req)
}

func (u useCase) FindGiroByDate(req *request.FillterTransactionByDate) ([]entity.TransactionAccount, error) {
	return u.repo.GetGiroByDate(req)
}

func (u useCase) FindVirtualAccountByDate(req *request.FillterTransactionByDate) ([]entity.TransactionVirtualAccount, error) {
	return u.repo.GetVirtualAccountByDate(req)
}

func (u useCase) GetListAccess() ([]string, error) {
	return u.repo.GetListAccess()
}
func (u useCase) GetAllTransaction(offset, limit int) ([]entity.MasterAccount, error) {
	return u.repo.GetAllTransaction(offset, limit)
}
func (u useCase) AssignRole(roleId uint, userId string) error {
	return u.repo.AssignRole(roleId, userId)
}
func (u useCase) GetAllRoles(offset, limit int) ([]entity.Role, error) {
	return u.repo.GetAllRoles(offset, limit)
}
func (u useCase) CreateRole(req *entity.Role) error {
	return u.repo.CreateRole(req)
}

func (u useCase) CreateAccess(access *entity.Access) error {
	return u.repo.CreateAccess(access)
}

func (u useCase) GetAllUser(offset, limit int) ([]entity.User, error) {
	return u.repo.GetAllUser(offset, limit)
}

func (u useCase) UpdateAccess(req *entity.Access, id uint) error {
	return u.repo.UpdateAccess(req, id)
}

func (u useCase) GetRoleById(id string) (*entity.Role, error) {
	return u.repo.GetRoleById(id)
}

func (u useCase) GetAllAccessByRoleId(id string) ([]entity.Access, error) {
	return u.repo.GetAllAccessByRoleId(id)
}

func (u useCase) DeleteAccess(id uint) error {
	return u.repo.DeleteAccess(id)
}

func (u useCase) DeleteRole(id string) error {
	return u.repo.DeleteRole(id)
}

func (u useCase) UpdateRole(role *entity.Role, id uint) error {
	return u.repo.UpdateRole(role, id)
}
func (u useCase) UserApprove(user *entity.User) error {
	return u.repo.UserApprove(user)
}

func (u useCase) GetById(id uint) (*entity.User, error) {
	return u.repo.GetById(id)
}

func (u useCase) DeleteUser(id uint) error {
	// Periksa apakah pengguna dengan ID tersebut ada dalam sistem
	_, err := u.repo.GetById(id)
	if err != nil {
		return err
	}

	// Melakukan penghapusan pengguna dari repository
	return u.repo.DeleteUser(id)
}
