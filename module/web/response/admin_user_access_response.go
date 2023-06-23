package response

type RoleUserResponse struct {
	Status string           `json:"status"`
	Data   ItemRoleResponse `json:"data"`
}
type ItemRoleResponse struct {
	Role   string       `json:"role"`
	Access []ItemAccess `json:"access"`
}

type ItemAccess struct {
	Resource string `json:"resource"`
	CanRead  bool   `json:"can_read"`
	CanWrite bool   `json:"can_write"`
}
