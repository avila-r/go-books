package services

import (
	"github.com/avila-r/books/app/models"
	"github.com/avila-r/books/pkg/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	database *gorm.DB
}

func NewUserService() *UserService {
	conn := db.GetConnection()

	return &UserService{conn}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	var (
		users []models.User
	)

	result := s.database.Find(&users)

	return users, result.Error
}

func (s *UserService) CreateUser(r models.UserRequest) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	// TODO: Assert that email is unique.
	user := models.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: string(hash),
		Role:     "common",
	}

	result := s.database.Create(&user)

	return &user, result.Error
}
