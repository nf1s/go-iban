package main

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
		if isNum(string(_char)) {
			ibanInNumbers += _char
		} else {
			ibanInNumbers += string('A' - 1 + char)
		}
	}
	return ibanInNumbers
}
