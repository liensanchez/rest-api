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

/*
	minJ := 1
	maxJ := 86
	randomNum := rand.Intn(maxJ-minJ+1) + minJ

	return func(c *fiber.Ctx) error {
		fmt.Println(randomNum)
		rows, err := db.Query("SELECT frases_column FROM frases WHERE id = $1;", randomNum)
		if err != nil {
			return err
		}
		defer rows.Close()

		var response []Test

		for rows.Next() {
			var text string
			err := rows.Scan(&text)
			if err != nil {
				return err
			}
			response = append(response, Test{text})
		}

		return c.JSON(response)
	}


func GetFrasesHandler(db *sql.DB) fiber.Handler {


}*/
