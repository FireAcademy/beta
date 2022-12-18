package main

import "github.com/gofiber/fiber/v2"

const COST_PER_CREDIT = 420

func MakeSuccessResponse(r fiber.Map, results int) fiber.Map {
	r["success"] = true
	if results < 1 {
		results = 1
	}
	r["cost"] = results * COST_PER_CREDIT

	return r
}

func MakeErrorResponse(r fiber.Map, message string) fiber.Map {
	r["success"] = false
	r["message"] = message
	r["cost"] = COST_PER_CREDIT

	return r
}