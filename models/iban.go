package models

import "database/sql"

type IbanFormat struct {
	CountryCode string `db:"countrycode"`
	Country     string `db:"country"`
	Size        string `db:"size"`
	BBANFormat  string `db:"bbanformat"`
	IBANFormat  string `db:"ibanformat"`
}

func GetIbanFormats(db *sql.DB, countryCode string) IbanFormat {
	var iban IbanFormat
	row := db.QueryRow("SELECT * FROM iban WHERE countrycode=$1;", countryCode)
	row.Scan(&iban.CountryCode, &iban.Country, &iban.Size, &iban.BBANFormat, &iban.IBANFormat)
	return iban
}
