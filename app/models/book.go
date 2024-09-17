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
	CompanyID   int64
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
	CompanyID   int64              `json:"company_id"`
	HasEbook    bool               `json:"has_ebook"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type BookRequest struct {
	Title       string  `json:"title" validate:"required,min=3,max=100"`        // Title is required, with a minimum of 3 and a maximum of 100 characters
	Available   bool    `json:"available"`                                      // No specific validation needed for booleans
	Categories  []int64 `json:"categories" validate:"required,dive,uuid"`       // List of UUIDs is required, each UUID must be valid
	Description string  `json:"description" validate:"required,min=10,max=500"` // Description is required, with a length between 10 and 500 characters
	Image       string  `json:"image" validate:"omitempty,url"`                 // Optional field, must be a valid URL if provided
	PDF         string  `json:"pdf" validate:"omitempty,url"`                   // Optional field, must be a valid URL if provided
	CompanyID   int64   `json:"company_id" validate:"required,uuid"`            // Required and must be a valid UUID for the company
	HasEbook    bool    `json:"has_ebook"`                                      // No specific validation needed for booleans
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
