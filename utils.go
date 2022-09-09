package main

import (
	"math/big"
	"regexp"
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

func strToInt(value string) *big.Int {
	intVal, _ := new(big.Int).SetString(value, 10)
	return intVal

}
