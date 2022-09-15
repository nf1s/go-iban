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

func (app *App) Initialize() {
	app.DB = initDB()
	app.Router = mux.NewRouter()
	DBMigrate()
}

func (app *App) Run() {
	app.Router.HandleFunc("/health", app.healthHandler).Methods("GET")
	app.Router.HandleFunc("/iban", app.ibanHandler).Methods("POST")
	fmt.Printf("server running at %s", PORT)
	err := http.ListenAndServe(PORT, app.Router)
	log.Fatal(err)
}
