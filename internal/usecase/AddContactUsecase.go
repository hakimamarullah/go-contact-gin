package usecase

import (
	"database/sql"

	"contact_chiv2/domain/contract"
	"contact_chiv2/domain/model"
)

type AddContactUseCase struct {
	CountryRepo contract.AddCountryRepoInterface
	AddressRepo contract.AddAddressRepoInterface
	PersonRepo  contract.AddPersonRepoInterface
	PhoneRepo   contract.AddPhoneRepoInterface
}

func NewAddContactUsecase(countryrepo contract.AddCountryRepoInterface, addressrepo contract.AddAddressRepoInterface, personrepo contract.AddPersonRepoInterface, phonerepo contract.AddPhoneRepoInterface) contract.AddContactUsecaseInterface {
	return &AddContactUseCase{
		CountryRepo: countryrepo,
		AddressRepo: addressrepo,
		PersonRepo:  personrepo,
		PhoneRepo:   phonerepo,
	}
}

func (usecase *AddContactUseCase) AddContact(data model.AddContactRequest) (lastinserted int64, err error) {
	var (
		txCountry    *sql.Tx = &sql.Tx{}
		txAddress    *sql.Tx = &sql.Tx{}
		txPerson     *sql.Tx = &sql.Tx{}
		txPhone      *sql.Tx = &sql.Tx{}
		transactions []*sql.Tx
	)

	defer func() {
		if err != nil {
			rollBackIfError(transactions)
		} else {
			go txCountry.Commit()
			go txAddress.Commit()
			go txPerson.Commit()
			go txPhone.Commit()
		}
	}()

	transactions = append(transactions, txCountry, txAddress, txPerson, txPhone)

	lastinserted, txCountry, err = usecase.CountryRepo.AddCountry(model.NewCountryData(data.CountryName, data.Region))
	if err != nil {
		return
	}

	lastinserted, txAddress, err = usecase.AddressRepo.AddAddress(model.NewAddressData(data.FullAddress, data.DistrictNumber, int(lastinserted)))
	if err != nil {
		return
	}

	lastinserted, txPerson, err = usecase.PersonRepo.AddPerson(model.NewPersonData(data.FirstName, data.LastName, data.Age, int(lastinserted)))
	if err != nil {
		return
	}

	lastinserted, txPhone, err = usecase.PhoneRepo.AddPhone(model.NewPhoneData(data.Number, data.IMEI, int(lastinserted)))

	return
}

func rollBackIfError(txs []*sql.Tx) {
	for _, v := range txs {
		v.Rollback()
	}
}
