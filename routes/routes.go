package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, db *sql.DB) {

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendString("Go to https://github.com/liensanchez/rest-api for documentation")
	})

	api := app.Group("/api")

	frases := api.Group("/frases")

	chistes := api.Group("/chistes")

	refranes := api.Group("/refranes")

	FrasesRoute(frases, db)

	ChistesRoute(chistes, db)

	RefranesRoute(refranes, db)

}
