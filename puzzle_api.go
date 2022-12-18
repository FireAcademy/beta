package main

import (
	"github.com/gofiber/fiber/v2"
)

type GetPuzzleArgs struct {
    Puzzle_hash string `json:"puzzle_hash" example:"6b665c0e059050f71a1c3e8a7d5b58e4e1d7abbd02d937e9b5ab5abfd7f8eaba"`   
} // @name GetPuzzleArgs

type _ struct {
	Success bool `json:"success" example:"true"`
	Results int `json:"results" example:"1"`
	Result SwaggerPuzzle `json:"result"`
} // @name GetPuzzleResponse

// GetPuzzle godoc
// @Summary returns the stored puzzle for a given puzzle_hash
// @Description Beta stores all revealed inner puzzles of singletons. Use this method to get it from the corresponding puzzle hash. The puzzle will be returned as a hex string.
// @Tags Puzzle
// @Accept json
// @Produce json
// @Param Body body GetPuzzleArgs true "The inner puzzle hash"
// @Success 200 {object} GetPuzzleResponse
// @Failure 401 {object} NoAPIKeyResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /get_puzzle [post]
func GetPuzzle(c *fiber.Ctx) error {
	args := new(GetPuzzleArgs)
	if err := c.BodyParser(args); err != nil {
		return MakeErrorResponse(c, err.Error())
	}

	if len(args.Puzzle_hash) != 64 {
		return MakeErrorResponse(c, "puzzle_hash is not 64 characters long :(")
	}

	var p Puzzle
	result := DB.First(&p, "puzzle_hash = ?", args.Puzzle_hash)
	if result.Error != nil {
		return MakeErrorResponse(c, result.Error.Error())
	}

	return MakeSuccessResponseForSingleObject(c, "result", PuzzleToJSON(p))
}

func SetupPuzzleAPIRoutes(app *fiber.App) { // more like route lol
	app.Post("/get_puzzle", GetPuzzle)
}