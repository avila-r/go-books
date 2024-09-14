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
	Categories  []Category `gorm:"many2many:book_categories;"` // can be empty
	Reviews     []Review   `gorm:"many2many:book_reviews;"`    // starts empty
	Description string
	Image       *string // nullable
	PDF         *string // nullable
	CompanyID   uuid.UUID
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
	CompanyID   uuid.UUID          `json:"company_id"`
	HasEbook    bool               `json:"has_ebook"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type BookRequest struct {
	Title       string      `json:"title"`
	Available   bool        `json:"available"`
	Categories  []uuid.UUID `json:"categories"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	PDF         string      `json:"pdf"`
	CompanyID   uuid.UUID   `json:"company_id"`
	HasEbook    bool        `json:"has_ebook"`
}

func (b *Book) Response() *BookResponse {
	categories_response := make([]CategoryResponse, len(b.Categories))

	for i, category := range b.Categories {
		categories_response[i] = category.ToResponse()
	}

	reviews_response := make([]ReviewResponse, len(b.Reviews))

	for i, review := range b.Reviews {
		reviews_response[i] = review.ToResponse()
	}

	return &BookResponse{
		Title:       b.Title,
		Available:   b.Available,
		Rating:      b.Rating,
		Categories:  categories_response,
		Reviews:     reviews_response,
		Description: b.Description,
		Image:       *b.Image,
		PDF:         *b.PDF,
		CompanyID:   b.CompanyID,
		HasEbook:    b.HasEbook,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}
