package handlers

import (
	"github.com/avila-r/books/app/services"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	s *services.BookService
}

func NewBookHandler() *BookHandler {
	s := services.NewBookService()

	return &BookHandler{s}
}

// ListBooks func gets all existent books.
// @Description Get all existent books.
// @Summary get all existent books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /v1/books [get]
func (*BookHandler) ListBooks(c *fiber.Ctx) error {
	return nil
}

// GetBookById func gets book by given ID or 404 error.
// @Description Get book by given ID.
// @Summary get book by given ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} models.Book
// @Router /v1/book/{id} [get]
func (*BookHandler) GetBookById(c *fiber.Ctx) error {
	return nil
}

// InsertBook func for creates a new book.
// @Description Create a new book.
// @Summary create a new book
// @Tags Book
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param book_attrs body models.BookAttrs true "Book attributes"
// @Success 200 {object} models.Book
// @Security ApiKeyAuth
// @Router /v1/book [post]
func (*BookHandler) InsertBook(c *fiber.Ctx) error {
	return nil
}

// UpdateBook func for updates book by given ID.
// @Description Update book.
// @Summary update book
// @Tags Book
// @Accept json
// @Produce json
// @Param id body string true "Book ID"
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param book_status body integer true "Book status"
// @Param book_attrs body models.BookAttrs true "Book attributes"
// @Success 201 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/book [put]
func (*BookHandler) UpdateBook(c *fiber.Ctx) error {
	return nil
}

// DeleteBook func for deletes book by given ID.
// @Description Delete book by given ID.
// @Summary delete book by given ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id body string true "Book ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/book [delete]
func (*BookHandler) DeleteBookById(c *fiber.Ctx) error {
	return nil
}
