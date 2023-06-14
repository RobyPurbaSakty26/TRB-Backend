package user

import (
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

func (c controller) create(req *web.UserCreateRequest) (*web.UserResponse, error) {
	user := entity.User{
		Fullname: req.Fullname,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err := c.useCase.create(&user)
	if err != nil {
		return nil, err
	}

	result := &web.UserResponse{
		Status: "Success",
		Data: web.ItemResponse{
			ID:       user.ID,
			Username: user.Username,
			Fullname: user.Fullname,
			Email:    user.Email,
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
