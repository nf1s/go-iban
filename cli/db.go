package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func dbSession() *sql.DB {

	var USER = os.Getenv("DB_USER")
	var PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_NAME = os.Getenv("DB_NAME")
	var DB_URL = fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", USER, PASSWORD, DB_NAME)
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
