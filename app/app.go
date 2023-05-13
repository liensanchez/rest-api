package app

import (
	"database/sql"
	"os"

	"github.com/gofiber/fiber/v2"
	"main.go/routes"
)

func StartServer(db *sql.DB) {

	PORT := os.Getenv("PORT")

	app := fiber.New()

	routes.Routes(app, db)

	app.Listen(PORT)
}
