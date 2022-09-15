package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

func initDB(dbURL string) *sql.DB {

	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func dbMigrate(dbURL string) {
	m, err := migrate.New(MIGRATIONS_DIR, dbURL)

	if err != nil {
		log.Fatal("db migrations failed")

	}
	m.Up()
}
