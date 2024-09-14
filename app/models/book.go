package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title       string
	Available   bool
	Rating      float32
	Categories  []Category
	Reviews     []Review
	Description string
	Image       string
	PDF         string
	CompanyID   string
	HasEbook    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookResponse struct {
	Title       string             `json:"title"`
	Available   bool               `json:"available"`
	Rating      float32            `json:"rating"`
	Categories  []CategoryResponse `json:"categories"`
	Reviews     []ReviewResponse   `json:"reviews"`
	Description string             `json:"description"`
	Image       string             `json:"image"`
	PDF         string             `json:"pdf"`
	CompanyID   string             `json:"company_id"`
	HasEbook    bool               `json:"has_ebook"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}
