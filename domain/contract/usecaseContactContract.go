package contract

import "contact_chiv2/domain/model"

type AddContactUsecaseInterface interface {
	AddContact(data model.AddContactRequest) (lastinserted int64, err error)
}

type GetContactUsecaseInterface interface {
	GetContactByNumber(data model.GetContactResponse)
	GetContactByIMEI(data model.GetContactResponse)
}

type GetAllContactUsecaseInterface interface {
	GetAllContact() (response []model.GetContactResponse, err error)
}
