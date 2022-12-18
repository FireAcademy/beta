package main

import (
	"github.com/gofiber/fiber/v2"
)

type GetPuzzleArgs struct {
    Puzzle_hash string `json: "puzzle_hash"`   
}

// GetPuzzle godoc
// @Description returns the stored puzzle for a given puzzle_hash
// @Summary Beta stores all revealed inner puzzles of singletons. Use this method to get it from the corresponding puzzle hash. The puzzle will be returned as a hex string.
// @Tags Puzzle
// @Accept json
// @Produce json
// @Param puzzle_hash body string true "Puzzle Hash"
// @Success 200 {object} SwaggerPuzzle
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