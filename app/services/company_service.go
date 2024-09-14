package services

import (
	"github.com/avila-r/books/app/models"
	"github.com/avila-r/books/pkg/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyService struct {
	database *gorm.DB
}

func NewCompanyService() *CompanyService {
	conn := db.GetConnection()

	return &CompanyService{conn}
}

func (s *CompanyService) GetCompanyByID(id uuid.UUID) (*models.Company, error) {
	return nil, nil
}
