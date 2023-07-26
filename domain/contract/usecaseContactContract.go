package contract

import "contact_chiv2/domain/model"

type AddContactUsecaseInterface interface {
	AddContact(data model.AddContactRequest) (lastinserted int64, err error)
}
