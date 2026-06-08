package main

import (
	"game-mini/internal/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
