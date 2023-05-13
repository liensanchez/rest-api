package services

import (
	"database/sql"
	"math/rand"

	"main.go/models"
)

func GetFrases(db *sql.DB) ([]models.Frases, error) {
	minJ := 1
	maxJ := 86
	randomNum := rand.Intn(maxJ-minJ+1) + minJ

	rows, err := db.Query("SELECT frases_column FROM frases WHERE id = $1;", randomNum)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var response []models.Frases

	for rows.Next() {
		var text string
		err := rows.Scan(&text)
		if err != nil {
			return nil, err
		}
		response = append(response, models.Frases{FraseGraciosa: text})
	}

	return response, nil
}
