package main

import (
	"github.com/gofiber/fiber/v2"
)

func SetupSyncAPIRoutes(app *fiber.App) {
	app.Get("/get_peak_synced_block", func(c *fiber.Ctx) error {
		return c.SendString("GET get_peak_synced_block")
	})
	app.Post("/get_peak_synced_block", func(c *fiber.Ctx) error {
		return c.SendString("POST get_peak_synced_block")
	})

	app.Post("/get_synced_block", func(c *fiber.Ctx) error {
		return c.SendString("POST get_last_synced_bloc")
	})

	app.Post("/get_synced_blocks", func(c *fiber.Ctx) error {
		return c.SendString("POST get_last_synced_block")
	})
}