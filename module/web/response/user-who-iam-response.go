package response

type WhoImResponse struct {
	Status string             `json:"status"`
	Data   WhoIamItemResponse `json:"data"`
}

type WhoIamItemResponse struct {
	ID         uint              `json:"id"`
	Username   string            `json:"username"`
	Email      string            `json:"email"`
	IsActive   bool              `json:"is_active"`
	RoleId     uint              `json:"role_id"`
	RoleName   string            `json:"role"`
	Permission []AccessItemWhoIm `json:"permission"`
}

type AccessItemWhoIm struct {
	Resource string `json:"resource"`
	CanRead  bool   `json:"can_read"`
	CanWrite bool   `json:"can_write"`
}
