package response

type ListRoleResponse struct {
	Status string     `json:"status"`
	Data   []ItemRole `json:"data"`
}
type ItemRole struct {
	Id     uint         `json:"id"`
	Name   string       `json:"name"`
	Access []AccessItem `json:"access"`
}

type AccessItem struct {
	Resource string `json:"resource"`
	CanRead  bool   `json:"can_read"`
	CanWrite bool   `json:"can_write"`
}
