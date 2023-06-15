package admin

import "trb-backend/module/web"

type controller struct {
	useCase UseCaseAdminInterface
}

type ControllerAdminInterface interface {
	getAllUser() (*web.AllUserResponse, error)
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
	result := &web.AllUserResponse{}

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
