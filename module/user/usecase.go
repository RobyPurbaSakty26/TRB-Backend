package user

import (
	"errors"
	"trb-backend/helpers"
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
	getUserAndRole(id uint) (*entity.User, error)
	updatePassword(user *entity.User, password string) error
	updateInputFalse(user *entity.User, count int) error
	updateIsActive(user *entity.User, isActive bool) error
	getById(id int) (*entity.User, error)
	getAccessByRoleId(id uint) ([]entity.Access, error)
	login(username string) (*entity.User, error)
	forgotPassword(email, usename, password string) error
}

func NewUseCase(repo UserRepositoryInterface) UseCaseInterface {
	return UseCase{repo: repo}
}

func (u UseCase) forgotPassword(email, username, password string) error {
	// check user by username
	user, err := u.repo.getByEmail(email)
	if err != nil || user.Username != username {
		return errors.New("user not found")
	}

	// has password
	newPassword, err := helpers.HashPass(password)
	if err != nil {
		return err
	}

	// update password
	err = u.repo.updatePassword(user, newPassword)
	if err != nil {
		return err
	}

	return nil
}

func (u UseCase) login(username string) (*entity.User, error) {
	user, err := u.repo.getByUsername(username)
	if err != nil {
		return nil, err
	}

	// reset input false if > 3 hour
	hourLastUpdate := helpers.IsThreeHours(user.UpdatedAt)
	if hourLastUpdate > 3 {
		err := u.repo.updateInputFalse(user, 0)
		if err != nil {
			return nil, err
		}
	}

	// check input false, if false > 3 status user false
	if user.InputFalse >= 3 {
		err = u.repo.updateStatusIsActive(user, false)
		if err != nil {
			return nil, err
		}
	}

	// get new data user
	user, _ = u.repo.getByUsername(username)

	// check is active
	if user.Active != true {
		return nil, errors.New("The account is not yet activated")
	}
	return user, nil
}

func (u UseCase) getAccessByRoleId(id uint) ([]entity.Access, error) {
	return u.repo.getAccessByRoleId(id)
}

func (u UseCase) create(user *entity.User) error {
	_, err := u.repo.getByEmail(user.Email)
	if err == nil {
		return errors.New("Email already registered")
	}

	_, err = u.repo.getByUsername(user.Username)
	if err == nil {
		return errors.New("Username already registered")
	}

	return u.repo.save(user)
}

func (u UseCase) getByEmail(email string) (*entity.User, error) {
	return u.repo.getByEmail(email)
}

func (u UseCase) getByUsername(username string) (*entity.User, error) {
	return u.repo.getByUsername(username)
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
