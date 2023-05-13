package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/handler"
)

func FrasesRoute(frases fiber.Router) {

	frases.Get("/", handler.GetFrasesHandler())
}
