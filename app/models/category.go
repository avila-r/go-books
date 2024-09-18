package models

type Category struct {
	ID          int64 `gorm:"primaryKey"`
	Title       string
	Description string
}

type CategoryResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (c *Category) ToResponse() CategoryResponse {
	return CategoryResponse{
		Title:       c.Title,
		Description: c.Description,
	}
}
