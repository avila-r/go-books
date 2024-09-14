package models

import (
	"github.com/google/uuid"
)

type Company struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string
	Description string
	Image       string
}

type CompanyResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
