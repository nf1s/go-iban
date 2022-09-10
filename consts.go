package main

import (
	"strconv"
)

const ALPHANUMERIC_CHARS = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getIbanToNumMap() map[string]string {
	m := make(map[string]string)
	for index, char := range ALPHANUMERIC_CHARS {
		m[string(char)] = strconv.Itoa(index)
	}
	return m
}

func getIbanRules() [][]string {
	return readCSVFile("data/ibans.csv")

}

var IBAN_NUM = getIbanToNumMap()
var IBAN_RULES = getIbanRules()

const ALPHA = "a"
const ALPHANUMERIC = "c"
const NUMERIC = "n"

var BBAN_TO_REGEX = map[string]string{ALPHA: "[A-Z]", ALPHANUMERIC: "[A-Za-z0-9]", NUMERIC: "[0-9]"}
