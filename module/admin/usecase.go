package admin

import "trb-backend/module/entity"

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
