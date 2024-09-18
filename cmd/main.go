package main

import (
	"errors"
	"log"
	"os"

	_ "github.com/avila-r/books/docs"
	"github.com/avila-r/books/pkg/routes"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

// @title Book Management
// @version 1.0
// @description This is a RESTful API built-in Golang
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	err_handler := func(c *fiber.Ctx, err error) error {
		c.Set(
			fiber.HeaderContentType,
			fiber.MIMETextPlainCharsetUTF8,
		)

		// Tries to parse 'error' as 'APIError'
		var e *fiber.Error

		if !errors.As(err, &e) {
			// Status: Internal Server Error [500]
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(e.Code).JSON(e)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: err_handler,
	})

	routes.Activate(app)

	url := os.Getenv("SERVER_URL")

	app.Listen(url)
}
