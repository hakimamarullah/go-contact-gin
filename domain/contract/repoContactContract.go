package contract

import (
	"database/sql"

	"contact_chiv2/domain/model"
)

// Interface yang mengimplementasikan repository yang berkaitan
// dengan penambahan data pada table country
type AddCountryRepoInterface interface {
	AddCountry(data model.Country) (lastinserted int64, tx *sql.Tx, err error)
}

type GetCountryRepoInterface interface {
	GetAllCountry() (country []model.Country, err error)
}

// Interface yang berkaitan terhadap DML pada table Address
type AddAddressRepoInterface interface {
	AddAddress(data model.Address) (lastinserted int64, tx *sql.Tx, err error)
}

// Interface yang berkaitan terhadap DML pada table Person
type AddPersonRepoInterface interface {
	AddPerson(data model.Person) (lastinserted int64, tx *sql.Tx, err error)
}

// Interface yang berkaitan terhadap DML pada table Phone
type AddPhoneRepoInterface interface {
	AddPhone(data model.Phone) (lastinserted int64, tx *sql.Tx, err error)
}
