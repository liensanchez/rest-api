package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"main.go/handler"
)

func ChistesRoute(frases fiber.Router, db *sql.DB) {

	frases.Get("/", handler.GetChistesHandler(db))
}
