package admin

import (
	"errors"
	"fmt"
	"strconv"
	"trb-backend/module/entity"
	"trb-backend/module/web/request"
	"trb-backend/module/web/response"
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
	getAllUser() (*response.AllUserResponse, error)
	getRoleUser(id string) (*response.RoleUserResponse, error)
	updateAccessUser(req *request.UpdateAccessRequest, id string) error
	UserApprove(id uint) (*response.UserApproveResponse, error)
	deleteUser(id uint) error
}

func NewAdminController(usecase UseCaseAdminInterface) ControllerAdminInterface {
	return controller{
		useCase: usecase,
	}
}

func (c controller) getAllUser() (*response.AllUserResponse, error) {
	users, err := c.useCase.getAllUser()
	if err != nil {
		return nil, err
	}
	result := &response.AllUserResponse{
		Status: "Success",
	}

	for _, user := range users {
		item := response.ItemResponse{
			ID:       user.ID,
			Username: user.Username,
			Fullname: user.Fullname,
			Email:    user.Email,
			IsActive: user.Active,
			Role:     user.Role.Name,
			RoleId:   user.RoleId,
		}
		result.Data = append(result.Data, item)
	}
	return result, nil
}

func (c controller) getRoleUser(id string) (*response.RoleUserResponse, error) {
	data, err := c.useCase.getUserWithRole(id)
	if err != nil {
		return nil, err
	}

	result := &response.RoleUserResponse{
		Status: "Success",
		Data: response.ItemRoleResponse{
			Fullname: data.Fullname,
			Role:     data.Role.Name,
		},
	}

	accesses, err := c.useCase.getAllAccessByRoleId(id)
	if err != nil {
		return nil, err
	}
	for _, access := range accesses {
		item := response.ItemAccess{
			Resource: access.Resource,
			CanRead:  access.CanRead,
			CanWrite: access.CanWrite,
		}
		result.Data.Access = append(result.Data.Access, item)
	}

	return result, nil
}

func (c controller) updateAccessUser(req *request.UpdateAccessRequest, id string) error {
	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("cannot parse id string to uint64")
	}
	idUint := uint(idUint64)
	fmt.Println("role name :", req.Role)
	role := &entity.Role{
		Name: req.Role,
	}
	err = c.useCase.updateRole(role, idUint)
	if err != nil {
		return err
	}

	data, err := c.useCase.getById(idUint)
	if err != nil {
		return err
	}
	err = c.useCase.userApprove(data)

	for _, access := range req.Data {
		accessReq := &entity.Access{
			Resource: access.Resource,
			CanRead:  access.CanRead,
			CanWrite: access.CanWrite,
		}
		err := c.useCase.updateAccess(accessReq, idUint)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c controller) UserApprove(id uint) (*response.UserApproveResponse, error) {

	data, err := c.useCase.getById(id)
	if err != nil {
		return nil, err
	}

	err = c.useCase.userApprove(data)

	data, _ = c.useCase.getById(id)

	res := &response.UserApproveResponse{
		Status: "Success",
		Data: response.UserApproveItems{
			ID:       data.ID,
			Fullname: data.Fullname,
			Username: data.Username,
			Email:    data.Email,
			IsActive: data.Active,
		},
	}
	return res, nil
}

func (c controller) deleteUser(id uint) error {
	// Cek apakah pengguna dengan ID tersebut ada dalam sistem
	user, err := c.useCase.getById(id)
	if err != nil {
		return err
	}

	// Hapus pengguna dari use case
	err = c.useCase.deleteUser(user.ID)
	if err != nil {
		return err
	}

	return nil
}
