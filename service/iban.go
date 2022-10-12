package service

import (
	"errors"
	"fmt"
	"iban/repository"
)

type IbanService interface {
	ValidateIban(iban string) (bool, error)
}

type ibanService struct {
	ibanRepository repository.IbanRepository
}

func NewIbanService(r repository.IbanRepository) IbanService {
	return ibanService{
		ibanRepository: r}
}

func (s ibanService) ValidateIban(iban string) (bool, error) {
	s.ibanRepository.SetIbanValue(iban)
	if !s.ibanRepository.IsAlphanumeric() {
		return false, errors.New("Iban is not alphanumeric")
	}
	if !s.ibanRepository.IsSizeCorrect() {
		size := s.ibanRepository.Size()
		requiredSize := s.ibanRepository.CountrySpecificIbanSize()
		msg := fmt.Sprintf("Iban size of %d is not correct, should be %d", size, requiredSize)
		return false, errors.New(msg)
	}

	if !s.ibanRepository.IsMod97() {
		return false, errors.New("mod 97 operation fails")
	}

	if !s.ibanRepository.IsBBANFormatCorrect() {
		msg := fmt.Sprintf("BBAN is not in the correct format, should be %s", s.ibanRepository.BBANFormat())
		return false, errors.New(msg)
	}

	return true, nil
}
