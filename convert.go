package main

import "github.com/gofiber/fiber/v2"

func SyncedBlockToJSON(sb SyncedBlock) fiber.Map {
	return fiber.Map{
		"header_hash": sb.HeaderHash,
		"height": sb.Height,
	}
}

func SingletonStateToJSON(ss SingletonState) fiber.Map {
	return fiber.Map{
		"coin_id": ss.CoinID,
		"melted": ss.Melted,
		"header_hash": ss.HeaderHash,
		"height": ss.Height,
		"parent_coin_id": ss.ParentCoinID,
		"puzzle_hash": ss.PuzzleHash,
		"amount": ss.Amount,
		"launcher_id": ss.LauncherID,
		"inner_puzzle_hash": ss.InnerPuzzleHash,
	}
}

func PuzzleToJSON(p Puzzle) fiber.Map {
	return fiber.Map{
		"puzzle_hash": p.PuzzleHash,
		"puzzle": p.Puzzle,
	}
}
