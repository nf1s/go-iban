package main

import "regexp"

type Payload struct {
	Iban string `json:iban`
}

func (p *Payload) isAlphanumeric() bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(p.Iban)

}
