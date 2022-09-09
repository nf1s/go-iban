package main

import (
	"math/big"
	"strings"
)

type Iban struct {
	Value string `json:"iban"`
}

func (iban *Iban) isAlphanumeric() bool {
	return isAlphanumeric(iban.Value)
}

func (iban *Iban) countryCode() string {
	return strings.ToUpper(iban.Value[:2])
}

func (iban *Iban) header() string {
	return iban.Value[:4]
}

func (iban *Iban) BBAN() string {
	return iban.Value[4:]
}

func (iban *Iban) invertedIban() string {
	return iban.BBAN() + iban.header()
}

func (iban *Iban) inNumbers() string {
	ibanInNumbers := ""
	for _, char := range iban.invertedIban() {
		_char := string(char)
		ibanInNumbers += IBAN_NUM[strings.ToUpper(_char)]
	}
	return ibanInNumbers
}

func (iban *Iban) mod97() int64 {
	mod97 := new(big.Int).Mod(strToInt(iban.inNumbers()), strToInt("97"))
	return mod97.Int64()
}

func (iban *Iban) isMod97() bool {
	return iban.mod97() == 1
}
