package main

import (
	"github.com/Erenalp06/otel-go/cmd"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(otelfiber.Middleware())
	
	cmd.InitializeApp(app)
}
