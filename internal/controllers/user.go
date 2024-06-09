package controllers

import (
	"github.com/Erenalp06/otel-go/internal/services"
	"github.com/Erenalp06/otel-go/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		UserService: service,
	}
}

func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not retrieve users"})
	}
	return c.JSON(users)
}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "request body could not be parsed"})
	}

	user.ID = 0

	createdUser, err := uc.UserService.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}

func (uc *UserController) GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid User ID"})
	}

	user, err := uc.UserService.GetUserById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	return c.JSON(user)
}

func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "request body could not be parsed"})
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid User ID"})
	}
	user.ID = uint(id)

	_, err = uc.UserService.GetUserById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	updatedUser, err := uc.UserService.UpdateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not update user"})
	}
	return c.JSON(updatedUser)
}

func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
    }

    _, err = uc.UserService.GetUserById(uint(id))
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
    }

    err = uc.UserService.DeleteUser(uint(id))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not delete user"})
    }

    return c.SendStatus(fiber.StatusNoContent)
}
