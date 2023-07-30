package usecase

import (
	"contact_ginv1/domain/contract"
	"contact_ginv1/domain/model"
)

type getAllContactUsecase struct {
	contactRepo contract.GetAllContactRepoInterface
}

func NewGetAllContactUsecase(repo contract.GetAllContactRepoInterface) contract.GetAllContactUsecaseInterface {
	return &getAllContactUsecase{
		contactRepo: repo,
	}
}

func (uc *getAllContactUsecase) GetAllContact() (response []model.GetContactResponse, err error) {
	res, err := uc.contactRepo.GetAllContact()
	if err != nil {
		return
	}

	var temp model.GetContactResponse
	for _, v := range res {
		temp.Number = v.Numbers.String
		temp.IMEI = v.IMEI.String
		temp.FirstName = v.FirstName.String
		temp.LastName = v.LastName.String
		temp.FullAddress = v.FullAddress.String
		temp.DistrictNumber = int(v.DistrictNumber.Int64)
		temp.CountryName = v.CountryName.String
		temp.Region = v.Region.String

		response = append(response, temp)
	}

	return
}
