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

	return MakeSuccessResponseForSingleObject(c, "synced_block", SyncedBlockToJSON(sb))
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

	return MakeSuccessResponseForSingleObject(c, "synced_block", SyncedBlockToJSON(sb))
}

type GetSyncedBlocksArgs struct {
	Start_height uint32 `json: "start"` 
	End_height uint32 `json: "end"`   
}

func getSyncedBlocks(c *fiber.Ctx) error {
	args := new(GetSyncedBlocksArgs)
	if err := c.BodyParser(args); err != nil {
		return MakeErrorResponse(c, err.Error())
	}

	if args.End_height <= args.Start_height {
		return MakeErrorResponse(c, "end_height less than or equal to start_height")
	}
	if args.End_height - args.Start_height > 100 {
		return MakeErrorResponse(c, "if you really need more than 100 blocks at a time, mesage us directly")
	}

	var synced_blocks []SyncedBlock
	result := DB.Order("height asc").Find(
		&synced_blocks,
		"height >= ? AND height < ?",
			args.Start_height, args.End_height,
	)
	if result.Error != nil {
		return MakeErrorResponse(c, result.Error.Error())
	}

	var synced_blocks_JSON []fiber.Map
	for _, synced_block := range synced_blocks {
		synced_blocks_JSON = append(synced_blocks_JSON, SyncedBlockToJSON(synced_block))
	}

	return MakeSuccessResponseForArray(c, "synced_blocks", synced_blocks_JSON)
}

func SetupSyncAPIRoutes(app *fiber.App) {
	app.Get("/get_peak_synced_block", getPeakSyncedBlock)
	app.Post("/get_peak_synced_block", getPeakSyncedBlock)

	app.Post("/get_synced_block", getSyncedBlock)

	app.Post("/get_synced_blocks", getSyncedBlocks)
}