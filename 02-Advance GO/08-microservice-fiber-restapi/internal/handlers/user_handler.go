package handlers

import (
	"08-microservice-fiber-restapi/internal/models"
	"08-microservice-fiber-restapi/internal/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user, error := h.service.GetUserByID(id)
	if error != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{"message": "User not found"})
	}
	return c.JSON(user)
}

func (h *UserHandler) GetUserByName(c *fiber.Ctx) error {
	name := c.Params("name")
	user, error := h.service.GetUserByName(name)
	if error != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{"message": "User not found"})
	}
	return c.JSON(user)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to get users"})
	}
	return c.JSON(users)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": "Invalid request body"})
	}
	if err := h.service.CreateUser(user); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"message": "Failed to create user"})
	}
	return c.JSON(user)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user, error := h.service.GetUserByID(id)
	if error != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{"message": "User not found"})
	}
	if user == nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{"message": "User not found"})
	}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": "Invalid request body"})
	}
	if err := h.service.UpdateUser(user); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"message": "Failed to update user"})
	}
	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user, error := h.service.GetUserByID(id)
	if error != nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{"message": "User not found"})
	}
	if user == nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{"message": "User not found"})
	}
	if err := h.service.DeleteUser(user); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"message": "Failed to delete user"})
	}
	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
