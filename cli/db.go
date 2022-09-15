package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func dbSession() *sql.DB {

	db, err := sql.Open("postgres", DB_URL)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

var db = dbSession()
