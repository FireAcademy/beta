package main

import "github.com/gofiber/fiber/v2"

const COST_PER_CREDIT = 420

func MakeSuccessResponseForSingleObject(c *fiber.Ctx, object_name string, object fiber.Map) error {
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"cost": COST_PER_CREDIT,
		object_name: object,
	})
}

func MakeSuccessResponseForArray(c *fiber.Ctx, array_name string, objects []fiber.Map) error {
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"cost": len(objects) * COST_PER_CREDIT,
		array_name: objects,
	})
}

func MakeErrorResponse(c *fiber.Ctx, message string) error {
	return c.Status(500).JSON(fiber.Map{
		"success": false,
		"message": message,
		"cost": COST_PER_CREDIT,
	})
}