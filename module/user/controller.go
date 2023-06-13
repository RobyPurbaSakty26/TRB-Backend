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
