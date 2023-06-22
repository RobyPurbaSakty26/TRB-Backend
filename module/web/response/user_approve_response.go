package response

type UserApproveResponse struct {
	Status string           `json:"status"`
	Data   UserApproveItems `json:"data"`
}
