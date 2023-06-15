package web

type AllUserResponse struct {
	Status string         `json:"status"`
	Data   []ItemResponse `json:"data"`
}
