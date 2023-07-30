package usecase

import (
	"contact_ginv1/domain/contract"
	"contact_ginv1/domain/model"
)

type getContactUsecase struct {
	countryRepo contract.GetCountryRepoInterface
	addressRepo contract.GetAddressRepoInterface
	personRepo  contract.GetPersonRepoInterface
	phoneRepo   contract.GetPhoneRepoInterface
}

func NewGetContactUsecase(countryrepo contract.GetCountryRepoInterface, addressrepo contract.GetAddressRepoInterface, personrepo contract.GetPersonRepoInterface, phonerepo contract.GetPhoneRepoInterface) contract.GetContactUsecaseInterface {
	return &getContactUsecase{
		countryRepo: countryrepo,
		addressRepo: addressrepo,
		personRepo:  personrepo,
		phoneRepo:   phonerepo,
	}
}

func (uc *getContactUsecase) GetContactByNumber(number string) (contact model.GetContactResponse, err error) {
	phone, err := uc.phoneRepo.GetPhoneByNumber(number)
	if err != nil {
		return
	}

	person, err := uc.personRepo.GetPersonById(phone.PersonId)
	if err != nil {
		return
	}

	address, err := uc.addressRepo.GetAddressById(person.AddressId)
	if err != nil {
		return
	}

	country, err := uc.countryRepo.GetCountryById(address.CountryId)
	if err != nil {
		return
	}
	contact = model.GetContactResponse{
		FirstName:      person.FirstName.String,
		LastName:       person.LastName.String,
		Age:            int(person.Age.Int64),
		FullAddress:    address.FullAddress.String,
		DistrictNumber: int(address.DistrictNumber.Int64),
		CountryName:    country.CountryName.String,
		Region:         country.Region.String,
		Number:         phone.Numbers.String,
		IMEI:           phone.IMEI.String,
	}
	return
}

func (uc *getContactUsecase) GetContactByIMEI(imei string) (contact model.GetContactResponse, err error) {
	phone, err := uc.phoneRepo.GetPhoneByIMEI(imei)
	if err != nil {
		return
	}

	person, err := uc.personRepo.GetPersonById(phone.PersonId)
	if err != nil {
		return
	}

	address, err := uc.addressRepo.GetAddressById(person.AddressId)
	if err != nil {
		return
	}

	country, err := uc.countryRepo.GetCountryById(address.CountryId)
	if err != nil {
		return
	}
	contact = model.GetContactResponse{
		FirstName:      person.FirstName.String,
		LastName:       person.LastName.String,
		Age:            int(person.Age.Int64),
		FullAddress:    address.FullAddress.String,
		DistrictNumber: int(address.DistrictNumber.Int64),
		CountryName:    country.CountryName.String,
		Region:         country.Region.String,
		Number:         phone.Numbers.String,
		IMEI:           phone.IMEI.String,
	}
	return
}
