package main

import (
	"github.com/gofiber/fiber/v2"
)

func getPeakSyncedBlock(c *fiber.Ctx) error {
	var sb SyncedBlock
	DB.Order("height desc").First(&sb)

	return c.JSON(
		MakeSuccessResponse(SyncedBlockToJSON(sb), 1),
	);
}

func SetupSyncAPIRoutes(app *fiber.App) {
	app.Get("/get_peak_synced_block", getPeakSyncedBlock)
	app.Post("/get_peak_synced_block", getPeakSyncedBlock)

	app.Post("/get_synced_block", func(c *fiber.Ctx) error {
		return c.SendString("POST get_last_synced_bloc")
	})

	app.Post("/get_synced_blocks", func(c *fiber.Ctx) error {
		return c.SendString("POST get_last_synced_block")
	})
}