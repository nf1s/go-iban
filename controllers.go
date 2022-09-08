package main

import (
	"encoding/json"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"Status": "Ok"})

}

func ibanHandler(w http.ResponseWriter, r *http.Request) {
	var payload Payload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	if !isAlphanumeric(payload.Iban) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		response(w, map[string]string{"ValidationError": "Iban is not alphanumeric"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"iban": payload.Iban})
}
