package test

import (
	"log"
	"time"

	"github.com/avila-r/books/app/models"
	"github.com/gofiber/fiber/v2"
)

func GetBook(c *fiber.Ctx) error {
	log.Println("Request received.")

	book := models.BookResponse{
		Title:     "Como treinar o seu dragão",
		Available: true,
		Rating:    4.6,

		Categories: []models.CategoryResponse{
			{
				Title:       "Terror",
				Description: "Lorem Ipsum is simply dummy text",
			},
		},
		Reviews: []models.ReviewResponse{
			{
				Name:    "João Silva",
				Rating:  4,
				Comment: "Ótimo livro!",
				Date:    time.Now(),
			},
		},

		Description: "Lorem Ipsum is simply dummy text...",
		Image:       "https://encadernacaocapadura.com/wp-content/uploads/2022/05/livro-capa-dura.jpg",
		PDF:         "https://eppg.fgv.br/sites/default/files/teste.pdf",
		CompanyID:   "202c0a95-d401-4598-82bc-4901f871d359",
		HasEbook:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return c.JSON(book)
}

type DatabaseCollection struct {
	Categories []models.CategoryResponse `json:"categories"`
	Books      []models.BookResponse     `json:"books"`
	Groups     []models.GroupResponse    `json:"groups"`
	Companies  []models.CompanyResponse  `json:"companies"`
	Users      []models.UserResponse     `json:"users"`
}

func ReturnDataCollection(c *fiber.Ctx) error {
	database := DatabaseCollection{
		Categories: []models.CategoryResponse{
			{
				Title:       "Terror",
				Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
			},
			{
				Title:       "Ficção",
				Description: "Lorem Ipsum has been the industry's standard dummy text ever since the 1500s.",
			},
		},

		Books: []models.BookResponse{
			{
				Title:       "Como treinar o seu dragão",
				Available:   true,
				Rating:      4.6,
				Categories:  []models.CategoryResponse{{Title: "Terror", Description: "Lorem Ipsum is simply dummy text"}},
				Reviews:     []models.ReviewResponse{{Name: "João Silva", Rating: 4, Comment: "Ótimo livro!", Date: time.Now()}},
				Description: "Um excelente livro para todas as idades.",
				Image:       "https://encadernacaocapadura.com/wp-content/uploads/2022/05/livro-capa-dura.jpg",
				PDF:         "https://eppg.fgv.br/sites/default/files/teste.pdf",
				CompanyID:   "202c0a95-d401-4598-82bc-4901f871d359",
				HasEbook:    true,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		},

		Groups: []models.GroupResponse{
			{
				Name: "Análise e Desenvolvimento de Sistemas (2024)",
				Members: []models.GroupMemberResponse{
					{Name: "usuario 3", Email: "usuario3@gmail.com"},
				},
			},
		},

		Companies: []models.CompanyResponse{
			{
				Name:        "Faculdade EAD SP",
				Description: "Lorem Ipsum is simply dummy text.",
				Image:       "https://img.freepik.com/vetores-gratis/vetor-de-gradiente-de-logotipo-colorido-de-passaro_343694-1365.jpg",
			},
		},

		Users: []models.UserResponse{
			{
				Name:  "João",
				Email: "jpedro@gmail.com",
				Roles: []string{"ADMIN"},
			},
		},
	}

	return c.JSON(database)
}
