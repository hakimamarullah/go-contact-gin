package usecase

import (
	"contact_chiv2/domain/contract"
	"contact_chiv2/domain/model"
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

func (uc *getContactUsecase) GetContactByNumber(data model.GetContactResponse) {
	return
}

func (uc *getContactUsecase) GetContactByIMEI(data model.GetContactResponse) {
	return
}
