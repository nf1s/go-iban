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

func (app *App) Initialize(user, pwd, dbname string) {
	app.DB = initDB(user, pwd, dbname)
	app.Router = mux.NewRouter()
}

func (app *App) Run() {
	app.Router.HandleFunc("/health", app.healthHandler).Methods("GET")
	app.Router.HandleFunc("/iban", app.ibanHandler).Methods("POST")
	fmt.Println("server running at 8080")
	err := http.ListenAndServe(":8080", app.Router)
	log.Fatal(err)
}
