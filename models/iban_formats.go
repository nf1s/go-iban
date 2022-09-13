package models

type IbanFormat struct {
	countryCode string `db:"countrycode"`
	country     string `db:"country"`
	size        string `db:"size"`
	BBANFormat  string `db:"bbanformat"`
	IBANFormat  string `db:"ibanformat"`
}

func getIbanFormat(countryCode string) IbanFormat {
	var iban IbanFormat
	row := db.QueryRow("SELECT * FROM iban WHERE countrycode=$1;", countryCode)
	row.Scan(&iban.countryCode, &iban.country, &iban.size, &iban.BBANFormat, &iban.IBANFormat)
	return iban
}
