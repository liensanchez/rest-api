package handler

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"main.go/services"
)

func GetFrasesHandler(db *sql.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {
		response, err := services.GetFrases(db)
		if err != nil {
			return err
		}
		return c.JSON(response)
	}
}

func GetChistesHandler(db *sql.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {
		response, err := services.GetChistes(db)
		if err != nil {
			return err
		}
		return c.JSON(response)
	}
}

func GetRefranesHandler(db *sql.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {
		response, err := services.GetRefranes(db)
		if err != nil {
			return err
		}
		return c.JSON(response)
	}
}
