package contract

import (
	"contact_chiv2/domain/model"
)

type GetCountryRepoInterface interface {
	GetAllCountry() (country []model.Country, err error)
}

// interface yang digunakan untuk mengambil data Addres
type GetAddressRepoInterface interface {
	GetAllAddress() (address []model.Address, err error)
	GetAddressById(id int) (address model.Address, err error)
}

// interface yang digunakan untuk mengambil data Person
type GetPersonRepoInterface interface {
	GetPersonById(id int) (person model.Person, err error)
	GetAllPerson() (person []model.Person, err error)
}

// interface yang digunakan untuk mengambil data Phone
type GetPhoneRepoInterface interface {
	GetPhoneById(id int) (phone model.Phone, err error)
	GetAllPhone() (person []model.Phone, err error)
}

// Interface yang mengimplementasikan repository yang berkaitan
// dengan penambahan data pada table country
type AddCountryRepoInterface interface {
	AddCountry(data model.Country) (lastinserted int64, err error)
}

// Interface yang berkaitan terhadap DML pada table Address
type AddAddressRepoInterface interface {
	AddAddress(data model.Address) (lastinserted int64, err error)
}

// Interface yang berkaitan terhadap DML pada table Person
type AddPersonRepoInterface interface {
	AddPerson(data model.Person) (lastinserted int64, err error)
}

// Interface yang berkaitan terhadap DML pada table Phone
type AddPhoneRepoInterface interface {
	AddPhone(data model.Phone) (lastinserted int64, err error)
}

type GetAllContactRepoInterface interface {
	GetAllContact() (contacts []model.Contact, err error)
}

type UpdatePhoneInterface interface{}

type TrxRepoInterface interface {
	StartTrx()
	DoneTrx(err error)
}
