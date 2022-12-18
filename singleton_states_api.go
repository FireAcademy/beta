package main

import (
	"github.com/gofiber/fiber/v2"
)

type GetSingletonStatesArgs struct {
	Coin_id string `json:"coin_id" format:"hex" example:"625799464319c8703ac2d0664af98cf45b9b306f7dcf717b1070d170bb5916a9"`
	Header_hash string `json:"header_hash" format:"hex" example:"796c33c3905150e649211fdd9ed42c7c418758c30c321271973a7c792a5bd403"`
	Height int64 `json:"height" format:"int64" example:"2174316"`
	Parent_coin_id string `json:"parent_coin_id" format:"hex" example:"e9676e8ce096c5be27dee2fbf2120054d206e4df2de9ef59c24a651d3c558c95"`
	Puzzle_hash string `json:"puzzle_hash" format:"hex" example:"0a5a9c760970ebcc094c6f9faa3d9730f066c7a8f7450841a94fc4fd59229bc2"`
	Amount int64 `json:"amount" format:"int64" example:"1"`
	Launcher_id string `json:"launcher_id" format:"hex" example:"f4dd6f4ec490974f7eb98223748f47340a9e9363b4c2dccc1932cdbbc54d03fd"`
	Inner_puzzle_hash string `json:"inner_puzzle_hash" format:"hex" example:"6b665c0e059050f71a1c3e8a7d5b58e4e1d7abbd02d937e9b5ab5abfd7f8eaba"`

	Limit int `json:"limit" example:"7"`
	Order_by string `json:"order_by" example:"amount"`
	Order string `json:"order" example:"desc"`
	Offset int `json:"offset" example:"0"`
}// @name GetSingletonStatesArgs

type _ struct {
	Success bool `json:"success" example:"true"`
	Results int `json:"results" example:"1"`
	SingletonStates []SingletonState `json:"singleton_states"`
} // @name GetSingletonStatesResponse

// GetSingletonStates godoc
// @Summary returns singleton states for given conditions
// @Description Singletons are like souls - when a coin dies (gets spent), they might move on to a new coin or disappear (melt). Beta views these transitions as changes of state.
// @Description This endpoint takes multiple optional arguments and returns the states that match the criteria. Most body parameters are self-explanatory: if they match a field of the returned struct model (e.g., 'height' or 'launcher_id'), they act as filters. Only states with the specified values will be returned.
// @Description The first special parameter is 'limit' - by default, this function returns 100 results at most. Use this parameter to tweak this value.
// @Description 'order_by' can be 'coin_id', 'header_hash', 'height', 'parent_coin_id', 'puzzle_hash', 'amount', 'launcher_id', or 'inner_puzzle_hash'. The default ordering is ascending, but that can be changed by setting the 'order' parameter to 'desc'.
// @Description If your query returns more than 100 results and you need all of them for some reasons, you can also use the 'offset' parameter.
// @Description Note: Two singleton states CAN have the same coin id but a different 'melted' value - the primary key is a composite one: (coin_id, melted)
// @Tags Singleton States
// @Accept json
// @Produce json
// @Param Body body GetSingletonStatesArgs true "The arguments (see description for more details)"
// @Success 200 {object} GetSingletonStatesResponse
// @Failure 401 {object} NoAPIKeyResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /get_singleton_states [post]
func GetSingletonStates(c *fiber.Ctx) error {
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
	app.Post("/get_singleton_states", GetSingletonStates)
}