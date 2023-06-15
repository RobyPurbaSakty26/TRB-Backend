package admin

import "trb-backend/module/entity"

type useCase struct {
	repo AdminRepositoryInterface
}

type UseCaseAdminInterface interface {
	getAllUser() ([]entity.User, error)
	getAccessByRoleId(id uint) (*entity.Access, error)
	updateAccess(access *entity.Access) error
}

func NewUseCase(repo AdminRepositoryInterface) UseCaseAdminInterface {
	return useCase{
		repo: repo,
	}
}

func (u useCase) getAllUser() ([]entity.User, error) {
	return u.repo.getAllUser()
}

func (u useCase) getAccessByRoleId(id uint) (*entity.Access, error) {
	return u.repo.getAccessByRoleId(id)
}

func (u useCase) updateAccess(access *entity.Access) error {
	return u.repo.updateAccess(access)
}
