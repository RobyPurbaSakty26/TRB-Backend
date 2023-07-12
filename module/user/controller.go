package user

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
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
	// getByEmail(email string) (*response.UserResponse, error)
	// getByUsername(username string) (*response.UserResponse, error)
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

	_, err := c.useCase.getByEmail(req.Email)
	if err == nil {
		return nil, errors.New("Email already registered")
	}

	_, err = c.useCase.getByUsername(req.Username)
	if err == nil {
		return nil, errors.New("Username already registered")
	}

	hashPass, _ := helpers.HashPass(req.Password)

	user := entity.User{
		Fullname: req.Fullname,
		Username: req.Username,
		Email:    req.Email,
		Password: hashPass,
	}
	err = c.useCase.create(&user)
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

// func (c controller) getByEmail(email string) (*response.UserResponse, error) {
// 	data, err := c.useCase.getByEmail(email)

// 	if err != nil {
// 		return nil, err
// 	}

// 	res := &response.UserResponse{
// 		Status: "Success",
// 		Data: response.ItemResponse{
// 			ID:       data.ID,
// 			Username: data.Username,
// 			Fullname: data.Fullname,
// 			Email:    data.Email,
// 		},
// 	}
// 	return res, nil
// }

// func (c controller) getByUsername(username string) (*response.UserResponse, error) {
// 	data, err := c.useCase.getByUsername(username)

// 	if err != nil {
// 		return nil, err
// 	}

// 	res := &response.UserResponse{
// 		Status: "Success",
// 		Data: response.ItemResponse{
// 			ID:       data.ID,
// 			Fullname: data.Fullname,
// 			Username: data.Username,
// 			Email:    data.Email,
// 			IsActive: data.Active,
// 			Role:     data.Role.Name,
// 			RoleId:   data.RoleId,
// 		},
// 	}
// 	return res, nil
// }

func isThreeHours(update time.Time) float64 {
	lastUpdate := update
	current := time.Now()
	gap := current.Sub(lastUpdate)
	hour := gap.Hours()

	return hour

}

func (c controller) login(req *request.LoginRequest) (*response.LoginResponse, *int, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	data, err := c.useCase.getByUsername(req.Username)
	if err != nil {
		return nil, nil, err
	}

	hour := isThreeHours(data.UpdatedAt)
	user := &entity.User{
		Email:      data.Email,
		InputFalse: data.InputFalse,
	}

	if hour > 3 {
		err := c.useCase.updateInputFalse(user, 0)
		if err != nil {
			return nil, nil, err
		}
	}

	if data.InputFalse >= 3 {
		err = c.useCase.updateIsActive(user, false)
		if err != nil {
			return nil, nil, err
		}
	}

	data, _ = c.useCase.getByUsername(req.Username)

	if data.Active != true {
		return nil, nil, errors.New("The account is not yet activated")
	}

	err = helpers.ComparePass([]byte(data.Password), []byte(req.Password))
	if err != nil {
		err := c.useCase.updateInputFalse(user, data.InputFalse+1)

		if err != nil {
			return nil, nil, err
		}
		false := data.InputFalse + 1
		return nil, &false, errors.New("Password wrong")
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

	err = c.useCase.updateInputFalse(user, 0)
	if err != nil {
		return nil, nil, err
	}

	return res, nil, nil
}

func (c controller) updatePassword(req *request.UpdatePasswordRequest) (*response.UpdatePasswordResponse, error) {
	data, err := c.useCase.getByEmail(req.Email)
	pass := helpers.ValidatePass(req.Password)
	if !pass {
		return nil, errors.New("please choose a stronger password. Try a mix of letters, numbers, and symbols")
	}

	if err != nil || req.Username != data.Username {
		return nil, errors.New("user not found")
	}

	user := &entity.User{
		Password: data.Password,
		Email:    data.Email,
	}

	newPassword, err := helpers.HashPass(req.Password)
	if err != nil {
		return nil, err
	}

	err = c.useCase.updatePassword(user, newPassword)
	if err != nil {
		return nil, err
	}

	res := &response.UpdatePasswordResponse{
		Status:  "Success",
		Message: "Password changed successfully",
	}
	return res, nil
}
