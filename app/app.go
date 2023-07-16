package app

import (
	"database/sql"
	"os"

	"github.com/gofiber/fiber/v2"
	"main.go/routes"
)

func StartServer(db *sql.DB) {

	app := fiber.New()

	routes.Routes(app, db)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	app.Listen("0.0.0.0:" + PORT)
}
