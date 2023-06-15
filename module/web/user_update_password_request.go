package web

type UpdatePasswordRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
