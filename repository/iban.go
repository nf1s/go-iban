package repository

import (
	"database/sql"
	"fmt"
	"iban/models"
	"math/big"
	"strings"
)

type IbanRepository interface {
	SetIbanValue(value string)

	Size() int
	BBAN() string
	BBANRegex() string
	IbanFormat() string
	BBANFormat() string
	CountrySpecificIbanSize() int

	IsAlphanumeric() bool
	IsSizeCorrect() bool
	IsMod97() bool
	IsBBANFormatCorrect() bool
}

type ibanRepository struct {
	DB    *sql.DB
	value string
}

func NewIbanRepository(db *sql.DB) IbanRepository {
	return &ibanRepository{DB: db}
}

func (iban *ibanRepository) SetIbanValue(value string) {
	iban.value = value
}

func (iban *ibanRepository) countryCode() string {
	return strings.ToUpper(iban.value[:2])
}

func (iban *ibanRepository) header() string {
	return iban.value[:4]
}

func (iban *ibanRepository) BBAN() string {
	return iban.value[4:]
}

func (iban *ibanRepository) invertedIban() string {
	return iban.BBAN() + iban.header()
}

func (iban *ibanRepository) inNumbers() string {
	ibanInNumbers := ""
	for _, char := range iban.invertedIban() {
		_char := string(char)
		ibanInNumbers += IBAN_NUM[strings.ToUpper(_char)]
	}
	return ibanInNumbers
}

func (iban *ibanRepository) mod97() int64 {
	mod97 := new(big.Int).Mod(strToBigInt(iban.inNumbers()), strToBigInt("97"))
	return mod97.Int64()
}

func (iban *ibanRepository) countryRules() models.IbanFormat {
	return models.GetIbanFormats(iban.DB, iban.countryCode())
}

func (iban *ibanRepository) country() string {
	return iban.countryRules().Country
}

func (iban *ibanRepository) CountrySpecificIbanSize() int {
	return strtoInt(iban.countryRules().Size)
}

func (iban *ibanRepository) BBANFormat() string {
	return iban.countryRules().BBANFormat
}

func (iban *ibanRepository) IbanFormat() string {
	return iban.countryRules().IBANFormat
}

func (iban *ibanRepository) BBANRegex() string {
	bbanFormats := strings.Split(iban.BBANFormat(), "-")
	var regex string

	for _, f := range bbanFormats {
		_type := f[len(f)-1:]
		_size := f[:len(f)-1]
		regex += fmt.Sprintf("%s{%s}", BBAN_TO_REGEX[_type], _size)
	}
	return "^" + regex + "$"
}

func (iban *ibanRepository) Size() int {
	return len([]rune(iban.value))
}

func (iban *ibanRepository) IsAlphanumeric() bool {
	return isAlphanumeric(iban.value)
}

func (iban *ibanRepository) IsSizeCorrect() bool {
	return iban.Size() == iban.CountrySpecificIbanSize()
}

func (iban *ibanRepository) IsMod97() bool {
	return iban.mod97() == 1
}

func (iban *ibanRepository) IsBBANFormatCorrect() bool {
	return match(iban.BBANRegex(), iban.BBAN())
}
