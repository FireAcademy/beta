package main

import (
	"github.com/gofiber/fiber/v2"
)

type _ struct {
	Success bool `json:"success" example:"true"`
	Results int `json:"results" example:"1"`
	SyncedBlock SyncedBlock `json:"synced_block"`
} // @name GetPeakSyncedBlockResponse

// GetPeakSyncedBlock godoc
// @Summary returns the peak synced block
// @Description Beta constantly syncs to the blockchain by fetching the latest blocks. This function returns information about the latest processed block (the one with the biggest 'height' value).
// @Tags Sync
// @Accept json
// @Produce json
// @Success 200 {object} GetPeakSyncedBlockResponse
// @Failure 401 {object} NoAPIKeyResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /get_peak_synced_block [get]
// @Router /get_peak_synced_block [post]
func GetPeakSyncedBlock(c *fiber.Ctx) error {
	var sb SyncedBlock
	result := DB.Order("height desc").First(&sb)
	if result.Error != nil {
		return MakeErrorResponse(c, result.Error.Error())
	}

	return MakeSuccessResponseForSingleObject(c, "synced_block", SyncedBlockToJSON(sb))
}

type GetSyncedBlockArgs struct {
	Height uint32 `json:"height" example:"2000000"`   
} // @name GetSyncedBlockArgs

type _ struct {
	Success bool `json:"success" example:"true"`
	Results int `json:"results" example:"1"`
	SyncedBlock SyncedBlock `json:"synced_block"`
} // @name GetSyncedBlockResponse

// GetSyncedBlock godoc
// @Summary returns block for given height
// @Description Beta constantly syncs to the blockchain by fetching the latest blocks. This function returns information about the block at the given height.
// @Tags Sync
// @Accept json
// @Produce json
// @Param Body body GetSyncedBlockArgs true "The height"
// @Success 200 {object} GetSyncedBlockResponse
// @Failure 401 {object} NoAPIKeyResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /get_synced_block [post]
func GetSyncedBlock(c *fiber.Ctx) error {
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
	Start_height uint32 `json:"start" example:"2000000"` 
	End_height uint32 `json:"end" example:"2000001"`
}// @name GetSyncedBlocksArgs

type _ struct {
	Success bool `json:"success" example:"true"`
	Results int `json:"results" example:"1"`
	SyncedBlocks []SyncedBlock `json:"synced_blocks"`
} // @name GetSyncedBlocksResponse

// GetSyncedBlocks godoc
// @Summary returns blocks for given range
// @Description Beta constantly syncs to the blockchain by fetching the latest blocks. This function returns information about the blocks with height in [start, end).
// @Tags Sync
// @Accept json
// @Produce json
// @Param Body body GetSyncedBlocksArgs true "The start and end heights"
// @Success 200 {object} GetSyncedBlocksResponse
// @Failure 401 {object} NoAPIKeyResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /get_synced_blocks [post]
func GetSyncedBlocks(c *fiber.Ctx) error {
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
	app.Get("/get_peak_synced_block", GetPeakSyncedBlock)
	app.Post("/get_peak_synced_block", GetPeakSyncedBlock)

	app.Post("/get_synced_block", GetSyncedBlock)

	app.Post("/get_synced_blocks", GetSyncedBlocks)
}