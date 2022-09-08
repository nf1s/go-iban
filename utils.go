package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

func isAlphanumeric(str string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(str)
}

func response(w http.ResponseWriter, body any) {
	json.NewEncoder(w).Encode(body)
}
