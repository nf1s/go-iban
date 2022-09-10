package main

import (
	"fmt"
	"math/big"
	"strings"
)

type Iban struct {
	Value string `json:"iban"`
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
	mod97 := new(big.Int).Mod(strToBigInt(iban.inNumbers()), strToBigInt("97"))
	return mod97.Int64()
}

func (iban *Iban) countryRules() map[string]string {
	for _, elem := range IBAN_RULES {
		countryCode := elem[0]

		if countryCode == iban.countryCode() {
			return map[string]string{"country_code": elem[0],
				"country":     elem[1],
				"size":        elem[2],
				"bban_format": elem[3],
				"iban_format": elem[4],
			}
		}

	}
	return map[string]string{}
}

func (iban *Iban) country() string {
	return iban.countryRules()["country"]
}

func (iban *Iban) countrySpecificIbanSize() int {
	return strtoInt(iban.countryRules()["size"])
}

func (iban *Iban) BBANFormat() string {
	return iban.countryRules()["bban_format"]
}

func (iban *Iban) IbanFormat() string {
	return iban.countryRules()["iban_format"]
}

func (iban *Iban) BBANRegex() string {
	bbanFormats := strings.Split(iban.BBANFormat(), "-")
	var regex string

	for _, f := range bbanFormats {
		_type := f[len(f)-1:]
		_size := f[:len(f)-1]
		regex += fmt.Sprintf("%s{%s}", BBAN_TO_REGEX[_type], _size)
	}
	return "^" + regex + "$"
}

func (iban *Iban) size() int {
	return len([]rune(iban.Value))
}

func (iban *Iban) isAlphanumeric() bool {
	return isAlphanumeric(iban.Value)
}

func (iban *Iban) isSizeCorrect() bool {
	return iban.size() == iban.countrySpecificIbanSize()
}

func (iban *Iban) isMod97() bool {
	return iban.mod97() == 1
}

func (iban *Iban) isBBANFormatCorrect() bool {
	return match(iban.BBANRegex(), iban.BBAN())
}
