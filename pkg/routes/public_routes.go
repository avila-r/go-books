package routes

import (
	"github.com/avila-r/books/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	book_handler := handlers.NewBookHandler()

	books := app.Get("/api/v1/books")

	books.Get("/books", book_handler.GetAllBooks)
}
