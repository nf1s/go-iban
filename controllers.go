package main

import (
	"encoding/json"
	"fmt"
	models "iban-go/models"
	"net/http"
	"runtime"
)

func response(w http.ResponseWriter, statusCode int, body any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func createMessage(key string, value any) map[string]any {
	return map[string]any{key: value}

}

func (app *App) healthHandler(w http.ResponseWriter, r *http.Request) {
	response(w, http.StatusOK, createMessage("status", "Ok"))

}

func (app *App) ibanHandler(w http.ResponseWriter, r *http.Request) {
	runtime.Breakpoint()
	iban := models.Iban{DB: app.DB}

	err := json.NewDecoder(r.Body).Decode(&iban)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	if !iban.IsAlphanumeric() {
		msg := createMessage("ValidationError", "Iban is not alphanumeric")
		response(w, http.StatusUnprocessableEntity, msg)
		return
	}

	if !iban.IsSizeCorrect() {
		size := iban.Size()
		requiredSize := iban.CountrySpecificIbanSize()
		err := fmt.Sprintf("Iban size of %d is not correct, should be %d", size, requiredSize)
		msg := createMessage("ValidationError", err)
		response(w, http.StatusUnprocessableEntity, msg)
		return
	}

	if !iban.IsMod97() {
		msg := createMessage("ValidationError", "mod 97 operation fails")
		response(w, http.StatusUnprocessableEntity, msg)
		return
	}

	if !iban.IsBBANFormatCorrect() {
		err := fmt.Sprintf("BBAN is not in the correct format, should be %s", iban.BBANFormat())
		msg := createMessage("ValidationError", err)
		response(w, http.StatusUnprocessableEntity, msg)
		return
	}

	response(w, http.StatusOK, createMessage("Iban", iban.IsBBANFormatCorrect()))

}
