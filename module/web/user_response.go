package web

type UserResponse struct {
	Status string       `json:"status"`
	Data   ItemResponse `json:"data"`
}
type ItemResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
	Role     string `json:"role"`
}
