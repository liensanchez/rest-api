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

	port := os.Getenv("PORT")

	if port == "" {
		port = "3010"
	}

	app.Listen("0.0.0.0:" + port)
}
