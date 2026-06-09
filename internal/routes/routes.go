package routes

import (
	"game-mini/internal/models"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	// Grouping
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("Hello")
	})

	// Games
	listGame := []models.Game{
		{
			ID:          1,
			Name:        "Mobile Legends",
			Slug:        "mobile-legends",
			Description: "Mobile Legends top up",
			ImageURL:    "https://example.com/mobile-legends.png",
			IsActive:    true,
		},
		{
			ID:          2,
			Name:        "Genshin Impact",
			Slug:        "genshin-impact",
			Description: "Genshin Impact top up",
			ImageURL:    "https://example.com/genshin-impact.png",
			IsActive:    true,
		},
		{
			ID:          3,
			Name:        "Valorant",
			Slug:        "valorant",
			Description: "Valorant top up",
			ImageURL:    "https://example.com/valorant.png",
			IsActive:    true,
		},
	}
	games := v1.Group("/games")
	games.Get("", func(ctx fiber.Ctx) error {
		return ctx.JSON(listGame)
	})
	games.Get("/:id", func(ctx fiber.Ctx) error {
		id := fiber.Params[int](ctx, "id")
		for _, game := range listGame {
			if game.ID == id {
				return ctx.JSON(game)
			}
		}
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "game not found"})
	})
	games.Post("", func(ctx fiber.Ctx) error {
		game := new(models.Game)
		if err := ctx.Bind().All(game); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// Update listGame
		game.ID = len(listGame) + 1
		listGame = append(listGame, *game)

		return ctx.Status(201).JSON(game)
	})
}
