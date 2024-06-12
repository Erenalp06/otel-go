package controllers

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/Erenalp06/otel-go/internal/services"
	"github.com/Erenalp06/otel-go/pkg/models"
	"github.com/Erenalp06/otel-go/util"
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

func (uc *UserController) ExternalAPI(c *fiber.Ctx) error {
	api1Url := "https://api.thecatapi.com/v1/images/search"

	httpClient := util.NewHTTPClient()
	ctx := context.Background()

	resp1, err := httpClient.Get(ctx, api1Url)
	if err != nil {
		fmt.Println("Dış API'ye yapılan çağrıda hata:", err)
		return err
	}
	defer resp1.Body.Close()

	if resp1.StatusCode != http.StatusOK {
		fmt.Println("API'den OK olmayan HTTP durumu:", resp1.StatusCode)
		return c.Status(resp1.StatusCode).SendString("Non-OK HTTP status received from API")
	}

	body, err := io.ReadAll(resp1.Body)
	if err != nil {
		fmt.Println("Yanıt gövdesi okunurken hata:", err)
		return err
	}
	return c.Status(resp1.StatusCode).Send(body)
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

	go func() {
		api1Url := "https://api1.example.com/some-endpoint"
		req1, _ := http.NewRequest("GET", api1Url, nil)

		client1 := &http.Client{}
		resp1, err := client1.Do(req1)
		if err != nil {
			fmt.Println("Error calling external API 1:", err)
			return
		}
		defer resp1.Body.Close()

		if resp1.StatusCode != http.StatusOK {
			fmt.Println("Non-OK HTTP status from API 1:", resp1.StatusCode)
		}

		api2Url := "https://api2.example.com/another-endpoint"
		req2, _ := http.NewRequest("GET", api2Url, nil)

		client2 := &http.Client{}
		resp2, err := client2.Do(req2)
		if err != nil {
			fmt.Println("Error calling external API 2:", err)
			return
		}
		defer resp2.Body.Close()

		if resp2.StatusCode != http.StatusOK {
			fmt.Println("Non-OK HTTP status from API 2:", resp2.StatusCode)
		}
	}()

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
