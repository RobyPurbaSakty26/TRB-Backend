package user

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"trb-backend/helpers"
	"trb-backend/module/entity"
	"trb-backend/module/web/request"
	"trb-backend/module/web/response"

	"github.com/joho/godotenv"
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
	useCase UseCaseInterface
}

type ControllerUserInterface interface {
	create(req *request.UserCreateRequest) (*response.UserResponse, error)
	login(req *request.LoginRequest) (*response.LoginResponse, *int, error)
	updatePassword(req *request.UpdatePasswordRequest) (*response.UpdatePasswordResponse, error)
	whoIm(id int) (*response.WhoImResponse, error)
}

func NewController(usecase UseCaseInterface) ControllerUserInterface {
	return controller{
		useCase: usecase,
	}
}

func (c controller) whoIm(id int) (*response.WhoImResponse, error) {
	// get user
	user, err := c.useCase.getById(id)
	if err != nil {
		return nil, err
	}

	// get access
	role_id := user.RoleId
	fmt.Println(role_id)
	accesses, err := c.useCase.getAccessByRoleId(role_id)
	if err != nil {
		return nil, err
	}

	fmt.Println(user.Role.Name)

	// response who I am
	res := &response.WhoImResponse{
		Status: "Success",
		Data: response.WhoIamItemResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			IsActive: user.Active,
			RoleName: user.Role.Name,
			RoleId:   user.RoleId,
		},
	}
	for _, access := range accesses {
		items := response.AccessItemWhoIm{
			Resource: access.Resource,
			CanRead:  access.CanRead,
			CanWrite: access.CanWrite,
		}
		res.Data.Permission = append(res.Data.Permission, items)
	}

	return res, nil

}

func (c controller) create(req *request.UserCreateRequest) (*response.UserResponse, error) {
	pass := helpers.ValidatePass(req.Password)
	if !pass {
		return nil, errors.New("Please choose a stronger password. Try a mix of letters, numbers, and symbols")
	}

	hashPass, _ := helpers.HashPass(req.Password)

	user := entity.User{
		Fullname: req.Fullname,
		Username: req.Username,
		Email:    req.Email,
		Password: hashPass,
	}
	err := c.useCase.create(&user)
	if err != nil {
		return nil, err
	}

	data, _ := c.useCase.getUserAndRole(user.ID)
	result := &response.UserResponse{
		Status: "Success",
		Data: response.ItemResponse{
			ID:       data.ID,
			Username: data.Username,
			Fullname: data.Fullname,
			Email:    data.Email,
			IsActive: data.Active,
			Role:     data.Role.Name,
		},
	}
	return result, nil
}

func (c controller) login(req *request.LoginRequest) (*response.LoginResponse, *int, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	data, err := c.useCase.login(req.Username)
	if err != nil {
		return nil, nil, err
	}

	err = helpers.ComparePass([]byte(data.Password), []byte(req.Password))
	if err != nil {
		err := c.useCase.updateInputFalse(data, data.InputFalse+1)

		if err != nil {
			return nil, nil, err
		}
		false_input := data.InputFalse
		return nil, &false_input, errors.New("Password wrong")
	}

	secret := os.Getenv("SECRET_KEY")

	token, err := helpers.GenerateToken(
		strconv.FormatUint(uint64(data.ID), 10),
		data.Username,
		strconv.FormatUint(uint64(data.RoleId), 10),
		data.Role.Name,
		secret)
	if err != nil {
		return nil, nil, err
	}

	res := &response.LoginResponse{
		Status: "Success",
		Data: response.LoginItemsResponse{
			Token:    token,
			IsActive: data.Active,
		},
	}

	err = c.useCase.updateInputFalse(data, 0)
	if err != nil {
		return nil, nil, err
	}

	return res, nil, nil
}

func (c controller) updatePassword(req *request.UpdatePasswordRequest) (*response.UpdatePasswordResponse, error) {
	pass := helpers.ValidatePass(req.Password)
	if !pass {
		return nil, errors.New("please choose a stronger password. Try a mix of letters, numbers, and symbols")
	}

	err := c.useCase.forgotPassword(req.Email, req.Username, req.Password)

	if err != nil {
		return nil, err
	}

	res := &response.UpdatePasswordResponse{
		Status:  "Success",
		Message: "Password changed successfully",
	}
	return res, nil
}
