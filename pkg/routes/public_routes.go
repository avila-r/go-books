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

	router.activate()
}

func (r *PublicRouter) activate() {
	book_handler := handlers.NewBookHandler()
	user_handler := handlers.NewUserHandler()

	books := r.app.Group("/api/v1/books")
	users := r.app.Group("/api/v1/users")

	books.Get("/", book_handler.ListBooks)
	books.Post("/", book_handler.InsertBook)

	users.Get("/", user_handler.ListUsers)
	users.Post("/", user_handler.RegisterUser)
}
