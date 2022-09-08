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

func (iban *Iban) getCountryCode() string {
	return strings.ToUpper(iban.Value[:2])
}

func (iban *Iban) getHeader() string {
	return iban.Value[:4]
}

func (iban *Iban) getBBAN() string {
	return iban.Value[4:]
}
