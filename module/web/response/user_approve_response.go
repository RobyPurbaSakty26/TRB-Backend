package response

import "trb-backend/module/web"

type UserApproveResponse struct {
	Status string               `json:"status"`
	Data   web.UserApproveItems `json:"data"`
}
