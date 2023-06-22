package request

type UpdateAccessRequest struct {
	Role string          `json:"role"`
	Data []AccessRequest `json:"data"`
}

type AccessRequest struct {
	Resource string `json:"resource"`
	CanRead  bool   `json:"can_read"`
	CanWrite bool   `json:"can_write"`
}
