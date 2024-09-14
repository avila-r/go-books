package handlers

import "github.com/gofiber/fiber/v2"

type BookHandler struct{}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func (*BookHandler) GetAllBooks(c *fiber.Ctx) error {
	return nil
}
