package admin

import (
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
	getAllUser() ([]entity.User, error)
	updateAccess(req *entity.Access, id uint) error
	getAllAccessByRoleId(id string) ([]entity.Access, error)
	getUserWithRole(id string) (*entity.User, error)
	updateRole(role *entity.Role, id uint) error
	userApprove(user *entity.User) error
	getById(id int) (*entity.User, error)
}

func NewUseCase(repo AdminRepositoryInterface) UseCaseAdminInterface {
	return useCase{
		repo: repo,
	}
}

func (u useCase) getAllUser() ([]entity.User, error) {
	return u.repo.getAllUser()
}

func (u useCase) updateAccess(req *entity.Access, id uint) error {
	return u.repo.updateAccess(req, id)
}

func (u useCase) getUserWithRole(id string) (*entity.User, error) {
	return u.repo.getUserWithRole(id)
}

func (u useCase) getAllAccessByRoleId(id string) ([]entity.Access, error) {
	return u.repo.getAllAccessByRoleId(id)
}

func (u useCase) updateRole(role *entity.Role, id uint) error {
	return u.repo.updateRole(role, id)
}
func (u useCase) userApprove(user *entity.User) error {
	return u.repo.userApprove(user)
}

func (u useCase) getById(id int) (*entity.User, error) {
	return u.repo.getById(id)
}
