package main

import (
	"log"
	"os"

	_ "github.com/avila-r/books/docs"
	"github.com/avila-r/books/pkg/routes"
	"github.com/avila-r/books/test"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	"github.com/joho/godotenv"
)

// @title Book Management
// @version 1.0
// @description This is a RESTful API built-in Golang
func main() {
	environment()

	app := fiber.New()

	middleware(app)
	routing(app)

	url := os.Getenv("SERVER_URL")

	app.Listen(url)
}

func environment() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}
}

func routing(app *fiber.App) {
	app.Get("/docs", swagger.HandlerDefault)
	app.Get("/", test.ReturnDataCollection)
	app.Get("/book", test.GetBook)
	routes.ActivatePublicRouter(app)
}

func middleware(app *fiber.App) {
	app.Use(
		cors.New(),
		logger.New(),
		// middlewares.JsonWebToken(),
	)
}
