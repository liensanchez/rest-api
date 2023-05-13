package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, db *sql.DB) {

	api := app.Group("/api")

	frases := api.Group("/frases")

	FrasesRoute(frases, db)

}
