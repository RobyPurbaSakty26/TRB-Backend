package user

import "trb-backend/module/entity"

type UseCase struct {
	repo UserRepositoryInterface
}

type UseCaseInterface interface {
	create(user *entity.User) error
}

func NewUseCase(repo UserRepositoryInterface) UseCaseInterface {
	return UseCase{repo: repo}
}

func (u UseCase) create(user *entity.User) error {
	return u.repo.save(user)
}
