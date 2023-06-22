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
	updateIsActive(user *entity.User, isActive bool) error
	getById(id int) (*entity.User, error)
	deleteUser(id uint) error
	createAccess(access *entity.Access) error
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

func (u UseCase) updateIsActive(user *entity.User, isActive bool) error {
	return u.repo.updateStatusIsActive(user, isActive)
}

func (u UseCase) getById(id int) (*entity.User, error) {
	return u.repo.getById(id)
}

func (u UseCase) createAccess(access *entity.Access) error {
	return u.repo.createAccess(access)
}

func (u UseCase) deleteUser(id uint) error {
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
