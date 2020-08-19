package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Okta@qwerty@tcp(localhost:3306)/db_example")

	if err != nil {
		log.Fatal(err)
	}
	return db
}
