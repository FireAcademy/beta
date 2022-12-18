package main

import (
	"github.com/gofiber/fiber/v2"
)

func getPeakSyncedBlock(c *fiber.Ctx) error {
	var sb SyncedBlock
	result := DB.Order("height desc").First(&sb)
	if result.Error != nil {
		return MakeErrorResponse(c, result.Error.Error())
	}

	return MakeSuccessResponse(c, SyncedBlockToJSON(sb), 1)
}

type GetSyncedBlockArgs struct {
    Height uint32 `json: "height"`   
}

func getSyncedBlock(c *fiber.Ctx) error {
	args := new(GetSyncedBlockArgs)
	if err := c.BodyParser(args); err != nil {
        return MakeErrorResponse(c, err.Error())
    }

	var sb SyncedBlock
	result := DB.First(&sb, "height = ?", args.Height)
	if result.Error != nil {
		return MakeErrorResponse(c, result.Error.Error())
	}

	return MakeSuccessResponse(c, SyncedBlockToJSON(sb), 1)
}

func SetupSyncAPIRoutes(app *fiber.App) {
	app.Get("/get_peak_synced_block", getPeakSyncedBlock)
	app.Post("/get_peak_synced_block", getPeakSyncedBlock)

	app.Post("/get_synced_block", getSyncedBlock)

	app.Post("/get_synced_blocks", func(c *fiber.Ctx) error {
		return c.SendString("POST get_last_synced_block")
	})
}