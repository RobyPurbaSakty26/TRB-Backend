package response

type PaginateRole struct {
	Page       int        `json:"page"`
	Limit      int        `json:"limit"`
	Total      int        `json:"total"`
	TotalPages float64    `json:"total_pages"`
	Data       []ItemRole `json:"data"`
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
