package database

import (
	"database/sql"
	"fmt"
	"os"
)

func StartDatabase(force bool) *sql.DB {

	//variables para ingresar a SQL
	HOST := os.Getenv("HOST")
	PORT_DB := os.Getenv("PORT_DB")
	USER := os.Getenv("USER")
	PASSWORD := os.Getenv("PASSWORD")
	DBNAME := os.Getenv("DBNAME")

	if HOST == "" {
		HOST = "localhost"
	}
	if PORT_DB == "" {
		PORT_DB = "5432"
	}
	if USER == "" {
		USER = "postgres"
	}
	if PASSWORD == "" {
		PASSWORD = "admin2023"
	}
	if DBNAME == "" {
		DBNAME = "rest-api"
	}

	//ingresamos con las constantes
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		HOST, PORT_DB, USER, PASSWORD, DBNAME)
	db, err := sql.Open("postgres", postgresqlDbInfo)

	//cachamos el error y lo mostramos en consola
	if err != nil {
		panic(err)
	}

	/* 	defer db.Close() esto cierra la conexion!!!!!OJO CUIDADO PELIGRO!*/

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Established a successful connection!")

	//en casa de ser necesario se reinicia
	resetDatabase(db, force)

	//asignamos a la var Global, la conexion a database
	//Global = db

	return db

}

func resetDatabase(db *sql.DB, force bool) {
	if force {
		query := `DO $$ DECLARE
					r RECORD;
			BEGIN
					FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
									EXECUTE 'TRUNCATE TABLE ' || quote_ident(r.tablename) || ' RESTART IDENTITY CASCADE;';
					END LOOP;
			END $$;`

		_, _ = db.Exec(query)
	}
}
