package routes

import (
	"github.com/Erenalp06/otel-go/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userController *controllers.UserController) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	users := v1.Group("/users")

	users.Post("/", userController.CreateUser)
	users.Get("/", userController.GetAllUsers)
	users.Get("/:id", userController.GetUserById)
	users.Put("/:id", userController.UpdateUser)
	users.Delete("/:id", userController.DeleteUser)
}
