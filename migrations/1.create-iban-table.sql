CREATE TABLE iban (
	countryCode		char(2) NOT NULL,
	country	 varchar NOT NULL,
	size	int NOT NULL,
	BBANFormat	 varchar NOT NULL,
	IBANFormat	 varchar NOT NULL
);
ALTER TABLE countryCode ADD PRIMARY KEY (iban);
