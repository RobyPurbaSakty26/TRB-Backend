package user

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

type UseCase struct {
	repo UserRepositoryInterface
}

type UseCaseInterface interface {
	create(user *entity.User) error
	getByEmail(email string) (*entity.User, error)
	getByUsername(username string) (*entity.User, error)
	createRoleUser(role *entity.Role) error
	getUserAndRole(id uint) (*entity.User, error)
	updatePassword(user *entity.User, password string) error
	updateInputFalse(user *entity.User, count int) error
  getAllUsers() ([]*entity.User, error)
  updateIsActive(user *entity.User, isActive bool) error
	userApprove(user *entity.User) error
	getById(id int) (*entity.User, error)
}

func NewUseCase(repo UserRepositoryInterface) UseCaseInterface {
	return UseCase{repo: repo}
}

func (u UseCase) create(user *entity.User) error {
	return u.repo.save(user)
}

func (u UseCase) getByEmail(email string) (*entity.User, error) {
	return u.repo.getByEmail(email)
}

func (u UseCase) getByUsername(username string) (*entity.User, error) {
	return u.repo.getByUsername(username)
}

func (u UseCase) createRoleUser(role *entity.Role) error {
	return u.repo.createRole(role)
}
func (u UseCase) getUserAndRole(id uint) (*entity.User, error) {
	return u.repo.getUserAndRole(id)
}

func (u UseCase) updatePassword(user *entity.User, password string) error {
	return u.repo.updatePassword(user, password)
}

func (u UseCase) updateInputFalse(user *entity.User, count int) error {
	return u.repo.updateInputFalse(user, count)
}


func (u UseCase) getAllUser() ([]*entity.User, error) {
	return u.repo.getAllUsers()

func (u UseCase) updateIsActive(user *entity.User, isActive bool) error {
	return u.repo.updateStatusIsActive(user, isActive)
}

func (u UseCase) userApprove(user *entity.User) error {
	return u.repo.userApprove(user)
}

func (u UseCase) getById(id int) (*entity.User, error) {
	return u.repo.getById(id)

}
