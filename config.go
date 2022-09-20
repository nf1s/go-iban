package main

import (
	"fmt"
	"os"
	"strconv"
)

var MIGRATIONS_DIR = "file://migrations"
var PORT = ":8080"

func DBHost() string {
	docker, _ := strconv.ParseBool(os.Getenv("DOCKER"))
	if docker {
		return "db"

	}
	return "localhost"

}

func getDBURL(user, password, dbname string) string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, DBHost(), dbname)
}
