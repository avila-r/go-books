package middlewares

import (
	"github.com/avila-r/books/pkg/auth"

	middleware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JsonWebToken() func(*fiber.Ctx) error {
	config := auth.GetMiddlewareConfig()

	return middleware.New(config)
}
