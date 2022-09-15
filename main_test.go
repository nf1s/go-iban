package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
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
	a.DB.Exec("DELETE FROM test")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS test (
	countryCode		char(2) PRIMARY KEY NOT NULL,
	country	 			varchar NOT NULL,
	size					int NOT NULL,
	BBANFormat	 	varchar NOT NULL,
	IBANFormat	 	varchar NOT NULL
)`

func populateTable() {
	runtime.Breakpoint()
	_, err := a.DB.Exec("INSERT INTO test VALUES($1, $2, $3, $4, $5)",
		"AL", "Albania", "28", "8n-16c", "ALkk bbb s sss x cccc cccc cccc cccc")
	if err != nil {
		log.Printf("There was an error with insering data")
	}

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestEmptyTable(t *testing.T) {
	clearTable()
	var jsonStr = []byte(`{"iban":"AL35202111090000000001234567"}`)
	req, _ := http.NewRequest(http.MethodPost, "/iban", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)
	var m any
	json.Unmarshal(response.Body.Bytes(), &m)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}
