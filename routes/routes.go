package routes

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GATO request!")
	})

}
