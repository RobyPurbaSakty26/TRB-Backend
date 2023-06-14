package user

import (
	"errors"
	"regexp"
	"trb-backend/helpers"
	"trb-backend/module/entity"
	"trb-backend/module/web"
)

type controller struct {
	useCase UseCaseInterface
}

type ControllerUserInterface interface {
	create(req *web.UserCreateRequest) (*web.UserResponse, error)
	getByEmail(email string) (*web.UserResponse, error)
	getByUsername(username string) (*web.UserResponse, error)
}

func NewController(usecase UseCaseInterface) ControllerUserInterface {
	return controller{
		useCase: usecase,
	}
}
func validatePass(p string) bool {
	if len(p) < 8 {
		return false
	}
	match, _ := regexp.MatchString("[A-Z]", p)
	if !match {
		return false
	}

	match, _ = regexp.MatchString("[!@#$%^&*()_+{}|:\"<>?]", p)
	if !match {
		return false
	}

	match, _ = regexp.MatchString("[0-9]", p)
	if !match {
		return false
	}

	return true
}

func (c controller) create(req *web.UserCreateRequest) (*web.UserResponse, error) {
	pass := validatePass(req.Password)
	if !pass {
		return nil, errors.New("Please choose a stronger password. Try a mix of letters, numbers, and symbols")
	}

	//_, err := c.useCase.getByEmail(req.Email)
	//if err == nil {
	//	return nil, err
	//}
	//
	//_, err = c.useCase.getByUsername(req.Username)
	//if err == nil {
	//	return nil, err
	//}

	hashPass, _ := helpers.HashPass(req.Password)

	role := entity.Role{
		Name: "user",
	}
	err := c.useCase.createRoleUser(&role)
	if err != nil {
		return nil, err
	}
	user := entity.User{
		Fullname: req.Fullname,
		Username: req.Username,
		Email:    req.Email,
		Password: hashPass,
		RoleId:   role.ID,
	}
	err = c.useCase.create(&user)
	if err != nil {
		return nil, err
	}

	data, _ := c.useCase.getUserAndRole(user.ID)
	result := &web.UserResponse{
		Status: "Success",
		Data: web.ItemResponse{
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

func (c controller) getByEmail(email string) (*web.UserResponse, error) {
	data, err := c.useCase.getByEmail(email)

	if err != nil {
		return nil, err
	}

	res := &web.UserResponse{
		Status: "Success",
		Data: web.ItemResponse{
			ID:       data.ID,
			Username: data.Username,
			Fullname: data.Fullname,
			Email:    data.Email,
		},
	}
	return res, nil
}

func (c controller) getByUsername(username string) (*web.UserResponse, error) {
	data, err := c.useCase.getByUsername(username)

	if err != nil {
		return nil, err
	}

	res := &web.UserResponse{
		Status: "Success",
		Data: web.ItemResponse{
			ID:       data.ID,
			Fullname: data.Fullname,
			Username: data.Username,
			Email:    data.Email,
		},
	}
	return res, nil
}
