package models

type Category struct {
	ID          int64 `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
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
