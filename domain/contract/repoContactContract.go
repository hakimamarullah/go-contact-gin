package contract

import (
	"contact_ginv1/domain/model"
)

type GetCountryRepoInterface interface {
	GetAllCountry() (country []model.Country, err error)
	GetCountryById(id int) (country model.Country, err error)
}

// GetAddressRepoInterface interface yang digunakan untuk mengambil data Address
type GetAddressRepoInterface interface {
	GetAllAddress() (address []model.Address, err error)
	GetAddressById(id int) (address model.Address, err error)
}

// GetPersonRepoInterface interface yang digunakan untuk mengambil data Person
type GetPersonRepoInterface interface {
	GetPersonById(id int) (person model.Person, err error)
	GetAllPerson() (person []model.Person, err error)
}

// GetPhoneRepoInterface interface yang digunakan untuk mengambil data Phone
type GetPhoneRepoInterface interface {
	GetPhoneById(id int) (phone model.Phone, err error)
	GetAllPhone() (person []model.Phone, err error)
	GetPhoneByNumber(number string) (phone model.Phone, err error)
	GetPhoneByIMEI(imei string) (phone model.Phone, err error)
}

// AddCountryRepoInterface Interface yang mengimplementasikan repository yang berkaitan
// dengan penambahan data pada table country
type AddCountryRepoInterface interface {
	AddCountry(data model.Country) (lastinserted int64, err error)
}

// AddAddressRepoInterface Interface yang berkaitan terhadap DML pada table Address
type AddAddressRepoInterface interface {
	AddAddress(data model.Address) (lastinserted int64, err error)
}

// AddPersonRepoInterface Interface yang berkaitan terhadap DML pada table Person
type AddPersonRepoInterface interface {
	AddPerson(data model.Person) (lastinserted int64, err error)
}

// AddPhoneRepoInterface Interface yang berkaitan terhadap DML pada table Phone
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
