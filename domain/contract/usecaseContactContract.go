package contract

import "contact_ginv1/domain/model"

type AddContactUsecaseInterface interface {
	AddContact(data model.AddContactRequest) (lastinserted int64, err error)
}

type GetContactUsecaseInterface interface {
	GetContactByNumber(number string) (contact model.GetContactResponse, err error)
	GetContactByIMEI(imei string) (contact model.GetContactResponse, err error)
}

type GetAllContactUsecaseInterface interface {
	GetAllContact() (response []model.GetContactResponse, err error)
}
