package handlers

import (
	"github.com/avila-r/books/app/models"
	"github.com/avila-r/books/app/services"
	"github.com/avila-r/books/pkg/validation"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler() *UserHandler {
	s := services.NewUserService()

	return &UserHandler{s}
}

// ListUsers func gets all registered users.
// @Description Get all registered users.
// @Summary get all registered users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /v1/users [get]
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	response, err := h.service.GetUsers()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(response)
}

// RegisterUser func for creates a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router api/v1/users [post]
func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var (
		r models.UserRequest
	)

	if err := c.BodyParser(&r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "unable to parse request",
		})
	}

	if err := validation.Validate(r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			validation.ValidatorErrors(err),
		)
	}

	result, err := h.service.CreateUser(r)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "was not possible to register user - " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}
