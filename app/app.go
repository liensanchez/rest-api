package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	fmt.Println("Hello world!")
	app := fiber.New()
	app.Listen(":3000")
}
