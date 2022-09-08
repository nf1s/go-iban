package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/iban", ibanHandler).Methods("POST")
	fmt.Println("server running at 8080")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)

}
