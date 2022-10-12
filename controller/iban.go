package controller

import (
	"encoding/json"
	"iban/service"
	"net/http"
)

type Iban struct {
	Value string `json:"iban"`
}

type IbanController interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
	ValidateIban(w http.ResponseWriter, r *http.Request)
}

type ibanController struct {
	ibanService service.IbanService
}

func NewIbanController(s service.IbanService) IbanController {
	return &ibanController{
		ibanService: s}
}

func response(w http.ResponseWriter, statusCode int, body any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func createMessage(key string, value any) map[string]any {
	return map[string]any{key: value}

}

func (c *ibanController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response(w, http.StatusOK, createMessage("status", "Ok"))

}

func (c *ibanController) ValidateIban(w http.ResponseWriter, r *http.Request) {
	var iban Iban

	err := json.NewDecoder(r.Body).Decode(&iban)
	if err != nil {
		response(w, http.StatusBadRequest, createMessage("Problem with request payload", err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	valid, err := c.ibanService.ValidateIban(iban.Value)

	if err != nil {
		response(w, http.StatusUnprocessableEntity, createMessage("Iban", map[string]any{"valid": valid, "error": err.Error()}))
		return
	}
	response(w, http.StatusOK, createMessage("Iban", map[string]bool{"valid": valid}))

}
