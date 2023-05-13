package handler

import "github.com/gofiber/fiber/v2"

func GetFrasesHandler() fiber.Handler {

	return func(c *fiber.Ctx) error {
		response := "This is a phrase for testing"

		return c.SendString(response)
	}
}
