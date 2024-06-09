package main

import (
	"github.com/Erenalp06/otel-go/cmd"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	cmd.InitializeApp(app)
}
