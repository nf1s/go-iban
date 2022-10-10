package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func DBHost() string {
	docker, _ := strconv.ParseBool(os.Getenv("DOCKER"))
	if docker {
		return "db"

	}
	return os.Getenv("DB_HOST")

}

func dbSession() *sql.DB {

	var USER = os.Getenv("DB_USER")
	var PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_NAME = os.Getenv("DB_NAME")
	var DB_URL = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", USER, PASSWORD, DBHost(), DB_NAME)
	db, err := sql.Open("postgres", DB_URL)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
