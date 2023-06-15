package user

import "trb-backend/module/entity"

type UseCase struct {
	repo UserRepositoryInterface
}

type UseCaseInterface interface {
	create(user *entity.User) error
	getByEmail(email string) (*entity.User, error)
	getByUsername(username string) (*entity.User, error)
	createRoleUser(role *entity.Role) error
	getUserAndRole(id uint) (*entity.User, error)
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
