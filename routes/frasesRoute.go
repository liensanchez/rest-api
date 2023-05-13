package routes

import "github.com/gofiber/fiber/v2"

func FrasesRoute(frases fiber.Router) {

	frases.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})
}
