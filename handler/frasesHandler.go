package handler

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type Test struct {
	Frase string `json:"frase"`
}

func GetFrasesHandler(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT text_column FROM frases")
		if err != nil {
			return err
		}
		defer rows.Close()

		var response []Test // replace YourStruct with the name of your struct

		for rows.Next() {
			var text string
			err := rows.Scan(&text)
			if err != nil {
				return err
			}
			response = append(response, Test{text}) // replace YourStruct with the name of your struct
		}

		return c.JSON(response)
	}
}

/*func GetFrasesHandler(db *sql.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {
		response, err :=
		if err != nil {
			return err
		}
		return c.JSON(response)
	}
}*/
