package web

type ItemUserRequest struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
	Role     uint   `json:"role"`
}
