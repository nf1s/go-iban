package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	dbURL := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", "test", "test", "test")
	a.Initialize(dbURL)
	ensureTableExists()
	populateTable()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM test;")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS test (
	countryCode		char(2) PRIMARY KEY NOT NULL,
	country	 			varchar NOT NULL,
	size					int NOT NULL,
	BBANFormat	 	varchar NOT NULL,
	IBANFormat	 	varchar NOT NULL
)`

func populateTable() {
	_, err := a.DB.Exec("INSERT INTO test VALUES($1, $2, $3, $4, $5);",
		"AL", "Albania", "28", "8n-16c", "ALkk bbb s sss x cccc cccc cccc cccc")
	if err != nil {
		log.Printf("There was an error with insering data")
	}

}

func executeRequest(method, path string, payload []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func checkResponseBody(t *testing.T, expected, actual any) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v. Got %v", expected, actual)
	}

}

func TestCorrectIban(t *testing.T) {
	clearTable()
	response := executeRequest(http.MethodPost, "/iban", []byte(`{"iban":"AL35202111090000000001234567"}`))
	checkResponseCode(t, http.StatusOK, response.Code)

	var body map[string]bool
	json.Unmarshal(response.Body.Bytes(), &body)
	expectedBody := map[string]bool{"Iban": true}
	checkResponseBody(t, expectedBody, body)
}

func TestInvalidIbanSize(t *testing.T) {
	clearTable()
	response := executeRequest(http.MethodPost, "/iban", []byte(`{"iban":"AL3520211109000000000123456"}`))
	checkResponseCode(t, http.StatusUnprocessableEntity, response.Code)

	var body map[string]string
	json.Unmarshal(response.Body.Bytes(), &body)
	expectedBody := map[string]string{"ValidationError": "Iban size of 27 is not correct, should be 28"}
	checkResponseBody(t, expectedBody, body)
}

func TestInvalidIbanMod97Operation(t *testing.T) {
	clearTable()
	response := executeRequest(http.MethodPost, "/iban", []byte(`{"iban":"ALAA202111090000000001234567"}`))
	checkResponseCode(t, http.StatusUnprocessableEntity, response.Code)

	var body map[string]string
	json.Unmarshal(response.Body.Bytes(), &body)
	expectedBody := map[string]string{"ValidationError": "mod 97 operation fails"}
	checkResponseBody(t, expectedBody, body)
}
