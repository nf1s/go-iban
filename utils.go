package main

import (
	"encoding/csv"
	"log"
	"math/big"
	"os"
	"regexp"
	"strconv"
)

func isAlphanumeric(value string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(value)

}

func isAlpha(value string) bool {
	return regexp.MustCompile(`^[a-zA-Z]*$`).MatchString(value)
}

func isNum(value string) bool {
	return regexp.MustCompile(`^[0-9]*$`).MatchString(value)
}

func strToBigInt(value string) *big.Int {
	intVal, _ := new(big.Int).SetString(value, 10)
	return intVal

}

func strtoInt(value string) int {
	_int, err := strconv.Atoi(value)
	if err != nil {
		panic("unable to convert ")

	}
	return _int

}

func readCSVFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func match(reg string, str string) bool {
	match, err := regexp.MatchString(reg, str)
	if err != nil {
		panic("BBAN is not valid")

	}
	return match

}
