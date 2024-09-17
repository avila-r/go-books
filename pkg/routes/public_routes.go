package routes

import (
	"github.com/avila-r/books/app/handlers"

	"github.com/gofiber/fiber/v2"
)

type PublicRouter struct {
	app *fiber.App
}

func ActivatePublicRouter(app *fiber.App) {
	router := &PublicRouter{app}

	router.bookRoutes()
}

func (r *PublicRouter) bookRoutes() {
	book_handler := handlers.NewBookHandler()

	books := r.app.Group("/api/v1/books")

	books.Get("/", book_handler.ListBooks)
	books.Post("/", book_handler.InsertBook)
}
