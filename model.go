package main

import "strings"

type Iban struct {
	Value string `json:"iban"`
}

func (iban *Iban) isAlphanumeric() bool {
	return isAlphanumeric(iban.Value)

}

func (iban *Iban) getIbanInNumbers() string {
	ibanInNumbers := ""
	for _, char := range iban.Value {
		_char := string(char)
		ibanInNumbers += IBAN_NUM[strings.ToUpper(_char)]
	}
	return ibanInNumbers
}
