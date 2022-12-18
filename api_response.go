package main

import "github.com/gofiber/fiber/v2"

func MakeSuccessResponseForSingleObject(c *fiber.Ctx, object_name string, object fiber.Map) error {
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"results": 1,
		object_name: object,
	})
}

func MakeSuccessResponseForArray(c *fiber.Ctx, array_name string, objects []fiber.Map) error {
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"results": len(objects),
		array_name: objects,
	})
}

func MakeErrorResponse(c *fiber.Ctx, message string) error {
	return c.Status(500).JSON(fiber.Map{
		"success": false,
		"message": message,
		"results": 0,
	})
}