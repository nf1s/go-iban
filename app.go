package main

import (
	"context"
	"database/sql"
	"fmt"
	"iban/controller"
	"iban/repository"
	"iban/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var ctx = context.Background()

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialize(dbURL string) {
	app.DB = initDB(dbURL)
	app.Router = mux.NewRouter()

	ibanRepository := repository.NewIbanRepository(app.DB)
	ibanService := service.NewIbanService(ibanRepository)
	ibanController := controller.NewIbanController(ibanService)

	app.Router.HandleFunc("/health", ibanController.HealthCheck).Methods("GET")
	app.Router.HandleFunc("/iban", ibanController.ValidateIban).Methods("POST")
}

func (app *App) Run() {
	fmt.Printf("server running at %s", PORT)
	err := http.ListenAndServe(PORT, app.Router)
	log.Fatal(err)
}
