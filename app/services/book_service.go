package services

import (
	`github.com/avila-r/books/app/models`
	"github.com/avila-r/books/pkg/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookService struct {
	database *gorm.DB
}

func NewBookService() *BookService {
	conn := db.GetConnection()

	return &BookService{conn}
}

func (s *BookService) GetBooks() {
}

func (s *BookService) GetBookById(id uuid.UUID) (*models.Book, error) {
	return nil, nil
}

func (s *BookService) CreateBook(b *models.Book) {
}

func (s *BookService) UpdateBook(b *models.Book) {
}

func (s *BookService) DeleteBookById(id uuid.UUID) {
}
