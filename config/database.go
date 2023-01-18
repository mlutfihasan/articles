package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() (db *sql.DB) {
	connStr := "postgres://postgres:postgrespw@localhost:32768/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db
}
