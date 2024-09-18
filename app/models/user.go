package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name     string
	Email    string
	Password string
	Role     string
}

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
