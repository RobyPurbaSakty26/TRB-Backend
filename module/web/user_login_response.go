package web

type LoginResponse struct {
	Status string             `json:"status"`
	Data   LoginItemsResponse `json:"data"`
}

type LoginItemsResponse struct {
	Token    string `json:"token"`
	IsActive bool   `json:"is_active"`
}
