package services

import (
	"database/sql"
	"math/rand"

	"main.go/models"
)

func GetFrases(db *sql.DB) ([]models.Result, error) {
	minJ := 1
	maxJ := 86
	randomNum := rand.Intn(maxJ-minJ+1) + minJ

	rows, err := db.Query("SELECT text_column FROM frases WHERE id = $1;", randomNum)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var response []models.Result

	for rows.Next() {
		var text string
		err := rows.Scan(&text)
		if err != nil {
			return nil, err
		}
		response = append(response, models.Result{FraseGraciosa: text})
	}

	return response, nil
}

func GetChistes(db *sql.DB) ([]models.Result, error) {
	minJ := 2
	maxJ := 247
	randomNum := rand.Intn(maxJ-minJ+1) + minJ

	rows, err := db.Query("SELECT text_column FROM chistes WHERE id = $1;", randomNum)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var response []models.Result

	for rows.Next() {
		var text string
		err := rows.Scan(&text)
		if err != nil {
			return nil, err
		}
		response = append(response, models.Result{FraseGraciosa: text})
	}

	return response, nil
}

func GetRefranes(db *sql.DB) ([]models.Result, error) {
	minJ := 1
	maxJ := 2784
	randomNum := rand.Intn(maxJ-minJ+1) + minJ

	rows, err := db.Query("SELECT text_column FROM refranes WHERE id = $1;", randomNum)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var response []models.Result

	for rows.Next() {
		var text string
		err := rows.Scan(&text)
		if err != nil {
			return nil, err
		}
		response = append(response, models.Result{FraseGraciosa: text})
	}

	return response, nil
}
