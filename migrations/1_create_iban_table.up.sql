BEGIN;

CREATE TABLE IF NOT EXISTS iban (
	countryCode		char(2) PRIMARY KEY NOT NULL,
	country	 			varchar NOT NULL,
	size					int NOT NULL,
	BBANFormat	 	varchar NOT NULL,
	IBANFormat	 	varchar NOT NULL
);

COMMIT;
