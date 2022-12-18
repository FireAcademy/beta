package main

import (
	"github.com/gofiber/fiber/v2"
)

type GetPuzzleArgs struct {
    Puzzle_hash string `json: "puzzle_hash"`   
}

func getPuzzle(c *fiber.Ctx) error {
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
	app.Post("/get_puzzle", getPuzzle)
}