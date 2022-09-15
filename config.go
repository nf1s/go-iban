package main

import (
	"fmt"
	"os"
)

var USER = os.Getenv("DB_USER")
var PASSWORD = os.Getenv("DB_PASSWORD")
var DB_NAME = os.Getenv("DB_NAME")
var DB_URL = fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", USER, PASSWORD, DB_NAME)
var MIGRATIONS_DIR = "file://migrations"
var PORT = ":8080"
