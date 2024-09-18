package models

type Company struct {
	ID          int64 `gorm:"primaryKey"`
	Name        string
	Description string
	Image       string
}

type CompanyResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
