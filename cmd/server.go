package cmd

import (
	"log"

	"github.com/Erenalp06/otel-go/internal/config"
	"github.com/Erenalp06/otel-go/internal/controllers"
	"github.com/Erenalp06/otel-go/internal/repository"
	"github.com/Erenalp06/otel-go/internal/services"
	"github.com/Erenalp06/otel-go/migrations"
	"github.com/Erenalp06/otel-go/pkg/database"
	"github.com/Erenalp06/otel-go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitializeApp(app *fiber.App) {

	config.InitTracer()

	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("Could not load the database")
	}

	err = migrations.MigrateUsers(db)
	if err != nil {
		log.Fatal("Could not migrate db")
	}

	userRepository := repository.NewRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowOrigins:     "*",
	}))

	//app.Use(adaptor.HTTPMiddleware(otelhttp.NewMiddleware("server")))
	routes.SetupRoutes(app, userController)

	app.Listen(":8085")

}
