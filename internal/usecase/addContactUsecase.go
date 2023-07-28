package usecase

import (
	"contact_chiv2/domain/contract"
	"contact_chiv2/domain/model"
)

type addContactUseCase struct {
	countryRepo contract.AddCountryRepoInterface
	addressRepo contract.AddAddressRepoInterface
	personRepo  contract.AddPersonRepoInterface
	phoneRepo   contract.AddPhoneRepoInterface
	dbTrx       contract.TrxRepoInterface
}

func NewAddContactUsecase(countryrepo contract.AddCountryRepoInterface, addressrepo contract.AddAddressRepoInterface, personrepo contract.AddPersonRepoInterface, phonerepo contract.AddPhoneRepoInterface, transaction contract.TrxRepoInterface) contract.AddContactUsecaseInterface {
	return &addContactUseCase{
		countryRepo: countryrepo,
		addressRepo: addressrepo,
		personRepo:  personrepo,
		phoneRepo:   phonerepo,
		dbTrx:       transaction,
	}
}

func (uc *addContactUseCase) AddContact(data model.AddContactRequest) (lastinserted int64, err error) {
	uc.dbTrx.StartTrx()

	defer func() {
		uc.dbTrx.DoneTrx(err)
	}()

	if err != nil {
		return
	}

	lastinserted, err = uc.countryRepo.AddCountry(model.NewCountryData(data.CountryName, data.Region))
	if err != nil {
		return
	}

	lastinserted, err = uc.addressRepo.AddAddress(model.NewAddressData(data.FullAddress, data.DistrictNumber, int(lastinserted)))
	if err != nil {
		return
	}

	lastinserted, err = uc.personRepo.AddPerson(model.NewPersonData(data.FirstName, data.LastName, data.Age, int(lastinserted)))
	if err != nil {
		return
	}

	lastinserted, err = uc.phoneRepo.AddPhone(model.NewPhoneData(data.Number, data.IMEI, int(lastinserted)))

	return
}
