package admin

import (
	"trb-backend/module/entity"
	"trb-backend/module/web"
)

type useCase struct {
	repo AdminRepositoryInterface
}

type UseCaseAdminInterface interface {
	getAllUser() ([]entity.User, error)
	updateAccess(access *entity.Access, req *web.AccessRequest, id uint) error
	getAllAccessByRoleId(id string) ([]entity.Access, error)
	getUserWithRole(id string) (*entity.User, error)
}

func NewUseCase(repo AdminRepositoryInterface) UseCaseAdminInterface {
	return useCase{
		repo: repo,
	}
}

func (u useCase) getAllUser() ([]entity.User, error) {
	return u.repo.getAllUser()
}

func (u useCase) updateAccess(access *entity.Access, req *web.AccessRequest, id uint) error {
	return u.repo.updateAccess(access, req, id)
}

func (u useCase) getUserWithRole(id string) (*entity.User, error) {
	return u.repo.getUserWithRole(id)
}

func (u useCase) getAllAccessByRoleId(id string) ([]entity.Access, error) {
	return u.repo.getAllAccessByRoleId(id)
}
