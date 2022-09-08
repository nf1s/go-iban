package main

import (
	"encoding/json"
	"net/http"
)

func response(w http.ResponseWriter, statusCode int, body any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func createMessage(key string, value string) map[string]string {
	return map[string]string{key: value}

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response(w, http.StatusOK, createMessage("status", "Ok"))

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
		msg := createMessage("ValidationError", "Iban is not alphanumeric")
		response(w, http.StatusUnprocessableEntity, msg)
		return
	}

	response(w, http.StatusOK, createMessage("Iban", iban.getIbanInNumbers()))

}
