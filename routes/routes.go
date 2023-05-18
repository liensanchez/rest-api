package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, db *sql.DB) {

	api := app.Group("/api")

	frases := api.Group("/frases")

	chistes := api.Group("/chistes")

	refranes := api.Group("/refranes")

	FrasesRoute(frases, db)

	ChistesRoute(chistes, db)

	RefranesRoute(refranes, db)

}
