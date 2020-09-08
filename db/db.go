package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectDatabase() *sql.DB {
	conexao := "user=postgres dbname=hangar_runners password=nop host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}
	return db
}
