package main

import (
	"log"
)

func populateIbans(countries [][]string) {
	db := dbSession()
	for index, elem := range countries {
		if index == 0 {
			continue
		}
		_, err := db.Exec("INSERT INTO iban VALUES($1, $2, $3, $4, $5)", elem[0], elem[1], elem[2], elem[3], elem[4])
		if err != nil {
			log.Printf("There was an error with insering %v", countries)
		}

	}

}
