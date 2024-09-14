package test

import (
	"github.com/avila-r/books/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetCategory(c *fiber.Ctx) error {
	category := models.CategoryResponse{
		Title:       "Título da Categoria",
		Description: "Descrição da Categoria",
	}

	return c.JSON(category)
}
