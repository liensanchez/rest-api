package app

import (
	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()
	app.Listen(":3000")
}
