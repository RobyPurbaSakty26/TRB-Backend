package response

type ListRoleResponse struct {
	Status string     `json:"status"`
	Data   []ItemRole `json:"data"`
}

type ItemRole struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
