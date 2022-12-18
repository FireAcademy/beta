package main

import "github.com/gofiber/fiber/v2"

const COST_PER_CREDIT = 420

func MakeSuccessResponse(c *fiber.Ctx, r fiber.Map, results int) error {
	r["success"] = true
	if results < 1 {
		results = 1
	}
	r["cost"] = results * COST_PER_CREDIT

	return c.Status(200).JSON(r)
}

func MakeErrorResponse(c *fiber.Ctx, message string) error {
	return c.Status(500).JSON(fiber.Map{
		"success": false,
		"message": message,
		"cost": COST_PER_CREDIT,
	})
}