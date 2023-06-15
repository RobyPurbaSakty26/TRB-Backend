package entity

import "gorm.io/gorm"

type Access struct {
	gorm.Model
	RoleId   uint   `json:"role_id"`
	Resource string `json:"resource"`
	CanRead  bool   `json:"can_read"`
	CanWrite bool   `json:"can_write"`
	Role     Role
}
