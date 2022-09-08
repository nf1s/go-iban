package main

import (
	"encoding/json"
	"net/http"
)

func response(w http.ResponseWriter, body any) {
	json.NewEncoder(w).Encode(body)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response(w, map[string]string{"status": "Ok"})

}

func ibanHandler(w http.ResponseWriter, r *http.Request) {
	var iban Iban

	err := json.NewDecoder(r.Body).Decode(&iban)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	if !iban.isAlphanumeric() {
		msg := "Iban is not alphanumeric"
		http.Error(w, msg, http.StatusUnprocessableEntity)
		return
	}

	response(w, map[string]string{"iban": iban.Value})

}
