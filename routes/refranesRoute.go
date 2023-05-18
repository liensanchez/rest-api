package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"main.go/handler"
)

func RefranesRoute(frases fiber.Router, db *sql.DB) {

	frases.Get("/", handler.GetRefranesHandler(db))
}
