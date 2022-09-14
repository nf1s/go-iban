package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func initDB(user, pwd, dbname string) *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", user, pwd, dbname))

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
