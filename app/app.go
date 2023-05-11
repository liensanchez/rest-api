package app

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func StartServer( /*db *sql.DB*/ ) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	PORT := os.Getenv("PORT")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})

	app.Listen(PORT)
}
