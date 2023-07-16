package admin

import (
	"errors"
	"trb-backend/module/entity"
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
	GetAllRoles(offset, limit int) ([]entity.Role, int64, error)
	CreateRole(req *entity.Role) error
	UpdateAccess(req *entity.Access, id uint) error
	GetAllAccessByRoleId(id uint) ([]entity.Access, error)
	GetRoleById(id uint) (*entity.Role, error)
	UpdateRole(role *entity.Role, id uint) error
	UserApprove(user *entity.User) error
	GetById(id uint) (*entity.User, error)
	DeleteUser(id uint) error
	CreateAccess(access *entity.Access) error
	DeleteRole(id uint) error
	AssignRole(roleId, userId uint) error
	GetAllTransaction(offset, limit int) ([]entity.MasterAccount, int64, error)
	GetListAccess() ([]string, error)
	findVirtualAccountByDate(accNo, startDate, endDate string) ([]entity.TransactionVirtualAccount, error)
	findGiroByDate(accNo, startDate, endDate string) ([]entity.TransactionAccount, error)
	TotalDataUser() (int64, error)
	findGiroByDatePagination(accNo, startDate, endDate string, limit, page int) ([]entity.TransactionAccount, error)
	findVaByDatePagination(accNo, startDate, endDate string, limit, page int) ([]entity.TransactionVirtualAccount, error)
	TotalDataTransactionGiro(accNo, startDate, endDate string) (int64, error)
	TotalDataTransactionVa(accNo, startDate, endDate string) (int64, error)
	totalGetUserByUsername(username string) (int64, error)
	totalGetUserByEmail(email string) (int64, error)
	getUserByUsername(email string, page, limit int) ([]entity.User, error)
	getUserByEmail(email string, page, limit int) ([]entity.User, error)
}

func NewUseCase(repo AdminRepositoryInterface) UseCaseAdminInterface {
	return useCase{
		repo: repo,
	}
}

func (u useCase) getUserByEmail(email string, page, limit int) ([]entity.User, error) {
	return u.repo.getUserByEmail(email, page, limit)
}

func (u useCase) getUserByUsername(username string, page, limit int) ([]entity.User, error) {
	return u.repo.getUserByUsername(username, page, limit)
}

func (u useCase) totalGetUserByEmail(email string) (int64, error) {
	return u.repo.totalGetUserByEmail(email)
}

func (u useCase) totalGetUserByUsername(username string) (int64, error) {
	return u.repo.totalGetUserByUsername(username)
}

func (u useCase) TotalDataTransactionGiro(accNo, startDate, endDate string) (int64, error) {
	return u.repo.TotalDataTransactionGiro(accNo, startDate, endDate)
}

func (u useCase) TotalDataTransactionVa(accNo, startDate, endDate string) (int64, error) {
	return u.repo.TotalDataTransactionGiro(accNo, startDate, endDate)
}

func (u useCase) TotalDataUser() (int64, error) {
	return u.repo.TotalDataUser()
}

func (u useCase) findGiroByDatePagination(accNo, startDate, endDate string, limit, page int) ([]entity.TransactionAccount, error) {
	return u.repo.getGiroByDatePagination(accNo, startDate, endDate, limit, page)
}

func (u useCase) findVaByDatePagination(accNo, startDate, endDate string, limit, page int) ([]entity.TransactionVirtualAccount, error) {
	return u.repo.getVaByDatePagination(accNo, startDate, endDate, limit, page)
}

func (u useCase) findGiroByDate(accNo, startDate, endDate string) ([]entity.TransactionAccount, error) {
	return u.repo.getGiroByDate(accNo, startDate, endDate)
}

func (u useCase) findVirtualAccountByDate(accNo, startDate, endDate string) ([]entity.TransactionVirtualAccount, error) {
	return u.repo.getVirtualAccountByDate(accNo, startDate, endDate)
}

func (u useCase) GetListAccess() ([]string, error) {
	return u.repo.GetListAccess()
}
func (u useCase) GetAllTransaction(offset, limit int) ([]entity.MasterAccount, int64, error) {
	count, err := u.repo.TotalDataMaster()
	if err != nil {
		return nil, 0, errors.New("failed get total data transaction")
	}
	data, err := u.repo.GetAllTransaction(offset, limit)
	if err != nil {
		return nil, 0, errors.New("failed get data transaction")
	}
	return data, count, nil
}
func (u useCase) AssignRole(roleId, userId uint) error {
	_, err := u.repo.GetById(userId)
	if err != nil {
		return errors.New("user id not found")
	}
	err = u.repo.AssignRole(roleId, userId)
	if err != nil {
		return errors.New("failed assign role id")
	}
	return nil
}
func (u useCase) GetAllRoles(offset, limit int) ([]entity.Role, int64, error) {
	count, err := u.repo.TotalDataRole()
	if err != nil {
		return nil, 0, errors.New("failed get total data roles")
	}
	data, err := u.repo.GetAllRoles(offset, limit)
	if err != nil {
		return nil, 0, errors.New("failed get data roles")
	}
	return data, count, nil
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
func (u useCase) GetRoleById(id uint) (*entity.Role, error) {
	return u.repo.GetRoleById(id)
}

func (u useCase) GetAllAccessByRoleId(id uint) ([]entity.Access, error) {
	_, err := u.repo.GetRoleById(id)
	if err != nil {
		return nil, errors.New("role id not found")
	}
	return u.repo.GetAllAccessByRoleId(id)
}

func (u useCase) DeleteRole(id uint) error {
	_, err := u.repo.GetRoleById(id)
	if err != nil {
		return errors.New("role id not found")
	}
	err = u.repo.DeleteAccess(id)
	if err != nil {
		return errors.New("failed delete access")
	}
	err = u.repo.DeleteRole(id)
	if err != nil {
		return errors.New("failed delete role")
	}
	return nil
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
