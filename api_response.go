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

type _ struct {
	Success bool `json:"success" example:"false"`
	Results int `json:"results" example:"0"`
	Message string `json:"message" example:"example error message"`
} // @name ErrorResponse

func MakeErrorResponse(c *fiber.Ctx, message string) error {
	return c.Status(500).JSON(fiber.Map{
		"success": false,
		"message": message,
		"results": 0,
	})
}

type _ struct {
	Success bool `json:"success" example:"false"`
	Message string `json:"message" example:"No API key provided."`
} // @name NoAPIKeyResponse