package model

import (
	"database/sql"
)

type Country struct {
	CountryName sql.NullString
	Region      sql.NullString
}

type Address struct {
	FullAddress    sql.NullString
	DistrictNumber sql.NullInt64
	CountryId      int
}

type Person struct {
	FirstName sql.NullString
	LastName  sql.NullString
	Age       sql.NullInt64
	AddressId int
}

type Phone struct {
	Numbers  sql.NullString
	IMEI     sql.NullString
	PersonId int
}

type Contact struct {
	Country
	Address
	Person
	Phone
}

func NewCountryData(countryname string, region string) (country Country) {
	country.CountryName.Scan(countryname)
	country.Region.Scan(region)

	return
}

func NewAddressData(address string, districtnumber int, countryid int) (addr Address) {
	addr.CountryId = countryid
	addr.FullAddress.Scan(address)
	addr.DistrictNumber.Scan(districtnumber)

	return
}

func NewPersonData(firstname, lastname string, age, addresid int) (person Person) {
	person.AddressId = addresid
	person.FirstName.Scan(firstname)
	person.LastName.Scan(lastname)
	person.Age.Scan(age)

	return
}

func NewPhoneData(numbers, imei string, personid int) (phone Phone) {
	phone.PersonId = personid
	phone.Numbers.Scan(numbers)
	phone.IMEI.Scan(imei)

	return
}
