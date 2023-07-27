package usecase

import (
	"contact_chiv2/domain/contract"
	"contact_chiv2/domain/model"
)

type AddContactUseCase struct {
	CountryRepo contract.AddCountryRepoInterface
	AddressRepo contract.AddAddressRepoInterface
	PersonRepo  contract.AddPersonRepoInterface
	PhoneRepo   contract.AddPhoneRepoInterface
	DbTrx       contract.TrxRepoInterface
}

func NewAddContactUsecase(countryrepo contract.AddCountryRepoInterface, addressrepo contract.AddAddressRepoInterface, personrepo contract.AddPersonRepoInterface, phonerepo contract.AddPhoneRepoInterface, transaction contract.TrxRepoInterface) contract.AddContactUsecaseInterface {
	return &AddContactUseCase{
		CountryRepo: countryrepo,
		AddressRepo: addressrepo,
		PersonRepo:  personrepo,
		PhoneRepo:   phonerepo,
		DbTrx:       transaction,
	}
}

func (usecase *AddContactUseCase) AddContact(data model.AddContactRequest) (lastinserted int64, err error) {
	usecase.DbTrx.StartTrx()

	defer func() {
		usecase.DbTrx.DoneTrx(err)
	}()

	if err != nil {
		return
	}

	lastinserted, err = usecase.CountryRepo.AddCountry(model.NewCountryData(data.CountryName, data.Region))
	if err != nil {
		return
	}

	lastinserted, err = usecase.AddressRepo.AddAddress(model.NewAddressData(data.FullAddress, data.DistrictNumber, int(lastinserted)))
	if err != nil {
		return
	}

	lastinserted, err = usecase.PersonRepo.AddPerson(model.NewPersonData(data.FirstName, data.LastName, data.Age, int(lastinserted)))
	if err != nil {
		return
	}

	lastinserted, err = usecase.PhoneRepo.AddPhone(model.NewPhoneData(data.Number, data.IMEI, int(lastinserted)))

	return
}
