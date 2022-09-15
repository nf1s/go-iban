package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialize(dbURL string) {
	dbMigrate(dbURL)
	app.DB = initDB(dbURL)
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/health", app.healthHandler).Methods("GET")
	app.Router.HandleFunc("/iban", app.ibanHandler).Methods("POST")
}

func (app *App) Run() {
	fmt.Printf("server running at %s", PORT)
	err := http.ListenAndServe(PORT, app.Router)
	log.Fatal(err)
}
