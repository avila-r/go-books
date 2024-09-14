package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name  string
	Email string
	Roles []string `gorm:"type:json"`
}

type UserResponse struct {
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}
