package main

import (
	"github.com/gofiber/fiber/v2"
)

type GetSingletonStatesArgs struct {
	Coin_id string `json: "coin_id"`
	Header_hash string `json: "header_hash"`
	Height int64 `json: "height"`
	Parent_coin_id string `json: "parent_coin_id"`
	Puzzle_hash string `json: "puzzle_hash"`
	Amount int64 `json: "amount"`
	Launcher_id string `json: "launcher_id"`
	Inner_puzzle_hash string `json: "inner_puzzle_hash"`

	Limit int `json: "limit"`
	Order_by string `json: "order_by"`
	Order string `json: "order"`
	Offset int `json: "offset"`
}

func getSingletonStates(c *fiber.Ctx) error {
	args := new(GetSingletonStatesArgs)
	if err := c.BodyParser(args); err != nil {
		return MakeErrorResponse(c, err.Error())
	}

	limit := 100
	if args.Limit != 0 {
		if args.Limit < 1 || args.Limit > 100 {
			return MakeErrorResponse(c, "limit should be between 1 and 100")
		} 

		limit = args.Limit
	}
	q := DB.Limit(limit)

	if args.Coin_id != "" {
		if len(args.Coin_id) != 64 {
			return MakeErrorResponse(c, "coin_id is not 64 characters long :(")
		}

		q = q.Where(&SingletonState{CoinID: args.Coin_id})
	}

	if args.Header_hash != "" {
		if len(args.Header_hash) != 64 {
			return MakeErrorResponse(c, "header_hash is not 64 characters long :(")
		}

		q = q.Where(&SingletonState{HeaderHash: args.Header_hash})
	}

	if args.Height != 0 {
		if args.Height < 1 {
			return MakeErrorResponse(c, "height is negative :(")
		}

		q = q.Where(&SingletonState{Height: args.Height})
	}

	if args.Parent_coin_id != "" {
		if len(args.Parent_coin_id) != 64 {
			return MakeErrorResponse(c, "parent_coin_id is not 64 characters long :(")
		}

		q = q.Where(&SingletonState{ParentCoinID: args.Parent_coin_id})
	}

	if args.Puzzle_hash != "" {
		if len(args.Puzzle_hash) != 64 {
			return MakeErrorResponse(c, "puzzle_hash is not 64 characters long :(")
		}

		q = q.Where(&SingletonState{PuzzleHash: args.Puzzle_hash})
	}

	if args.Amount != 0 {
		if args.Amount < 1 {
			return MakeErrorResponse(c, "amount is negative :(")
		}

		q = q.Where(&SingletonState{Amount: args.Amount})
	}

	if args.Launcher_id != "" {
		if len(args.Launcher_id) != 64 {
			return MakeErrorResponse(c, "launcher_id is not 64 characters long :(")
		}

		q = q.Where(&SingletonState{LauncherID: args.Launcher_id})
	}

	if args.Inner_puzzle_hash != "" {
		if len(args.Launcher_id) != 64 {
			return MakeErrorResponse(c, "launcher_id is not 64 characters long :(")
		}

		q = q.Where(&SingletonState{InnerPuzzleHash: args.Inner_puzzle_hash})
	}

	if args.Offset != 0 {
		q = q.Offset(args.Offset)
	}

	if args.Order_by != "" {
		order := "asc"
		if args.Order == "desc" {
			order = "desc"
		}

		switch args.Order_by {
		case "coin_id",
			 "header_hash",
			 "height",
			 "parent_coin_id",
			 "puzzle_hash",
			 "amount",
			 "launcher_id",
			 "inner_puzzle_hash":
			q = q.Order(args.Order_by + " " + order)
		default:
			return MakeErrorResponse(c, "order_by is invalid")
		}
	}

	var singleton_states []SingletonState
	result := q.Find(&singleton_states)
	if result.Error != nil {
		return MakeErrorResponse(c, result.Error.Error())
	}

	var singleton_states_JSON []fiber.Map
	for _, singleton_state := range singleton_states {
		singleton_states_JSON = append(singleton_states_JSON, SingletonStateToJSON(singleton_state))
	}

	return MakeSuccessResponseForArray(c, "singleton_states", singleton_states_JSON)
}

func SetupSingletonStatesAPIRoutes(app *fiber.App) {
	app.Post("/get_singleton_states", getSingletonStates)
}