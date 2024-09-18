package routes

import (
	"github.com/avila-r/books/app/handlers"
	"github.com/avila-r/books/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/swagger"
)

type Router struct {
	app *fiber.App
}

func Activate(app *fiber.App) {
	router := &Router{app}

	router.enable()
}

func (r *Router) enable() {
	// Public route for API documentation
	r.app.Get("/docs", swagger.HandlerDefault)

	// Enable CORS and logging for the following routes
	r.app.Use(
		cors.New(),
		logger.New(),
	)

	// Protected route group
	api := r.app.Group("/api/v1")

	api.Use(
		middlewares.JsonWebToken(),
	)

	// Users routes
	user_handler := handlers.NewUserHandler()

	users := api.Group("/users")

	users.Get("/",
		user_handler.ListUsers,
	)

	// Books routes
	book_handler := handlers.NewBookHandler()

	books := api.Group("/books")

	books.Get("/",
		book_handler.ListBooks,
	)

	books.Post("/",
		book_handler.InsertBook,
	)

	// Public route group
	auth := r.app.Group("/auth")

	auth.Post("/register",
		user_handler.RegisterUser,
	)

	auth.Post("/login",
		func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"error": "login not yet implemented",
			})
		},
	)
}
