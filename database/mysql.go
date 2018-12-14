package database

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/parking_system")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
