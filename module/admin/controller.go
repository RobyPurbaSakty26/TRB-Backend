package admin

import (
	"strconv"
	"trb-backend/module/entity"
	"trb-backend/module/web"
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

type controller struct {
	useCase UseCaseAdminInterface
}

type ControllerAdminInterface interface {
	getAllUser() (*web.AllUserResponse, error)
	getRoleUser(id string) (*web.RoleUserResponse, error)
	updateAccessUser(req *web.UpdateAccessRequest, id string) error
}

func NewAdminController(usecase UseCaseAdminInterface) ControllerAdminInterface {
	return controller{
		useCase: usecase,
	}
}

func (c controller) getAllUser() (*web.AllUserResponse, error) {
	users, err := c.useCase.getAllUser()
	if err != nil {
		return nil, err
	}
	result := &web.AllUserResponse{
		Status: "Success",
	}

	for _, user := range users {
		item := web.ItemResponse{
			ID:       user.ID,
			Username: user.Username,
			Fullname: user.Fullname,
			Email:    user.Email,
			IsActive: user.Active,
			Role:     user.Role.Name,
		}
		result.Data = append(result.Data, item)
	}
	return result, nil
}

func (c controller) getRoleUser(id string) (*web.RoleUserResponse, error) {
	data, err := c.useCase.getUserWithRole(id)
	if err != nil {
		return nil, err
	}

	result := &web.RoleUserResponse{
		Status: "Success",
		Data: web.ItemRoleResponse{
			Fullname: data.Fullname,
			Role:     data.Role.Name,
		},
	}
	accesses, err := c.useCase.getAllAccessByRoleId(id)
	if err != nil {
		return nil, err
	}
	for _, access := range accesses {
		item := web.ItemAccess{
			Resource: access.Resource,
			CanRead:  access.CanRead,
			CanWrite: access.CanWrite,
		}
		result.Data.Access = append(result.Data.Access, item)
	}

	return result, nil
}

func (c controller) updateAccessUser(req *web.UpdateAccessRequest, id string) error {
	idUint64, _ := strconv.ParseUint(id, 10, 64)
	idUint := uint(idUint64)
	for _, access := range req.Data {
		accessReq := &entity.Access{
			Resource: access.Resource,
			CanRead:  access.CanRead,
			CanWrite: access.CanWrite,
		}
		err := c.useCase.updateAccess(accessReq, &access, idUint)
		if err != nil {
			return err
		}
	}
	return nil
}
