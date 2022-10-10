package main

import (
	"fmt"
	"os"
	"strconv"
)

var MIGRATIONS_DIR = "file://migrations"
var PORT = ":8000"

func DBHost() string {
	docker, _ := strconv.ParseBool(os.Getenv("DOCKER"))
	if docker {
		return "db"

	}
	return os.Getenv("DB_HOST")

}

func getDBURL(user, password, dbname string) string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, DBHost(), dbname)
}
