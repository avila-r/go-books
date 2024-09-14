package models

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID      uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name    string
	Rating  float32
	Comment string
	Date    time.Time
}

type ReviewResponse struct {
	Name    string    `json:"name"`
	Rating  float32   `json:"rating"`
	Comment string    `json:"comment"`
	Date    time.Time `json:"date"`
}
