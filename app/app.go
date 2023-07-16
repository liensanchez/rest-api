package app

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"main.go/routes"
)

func StartServer(db *sql.DB) {

	app := fiber.New()

	routes.Routes(app, db)

	app.Listen(":3010")
}
