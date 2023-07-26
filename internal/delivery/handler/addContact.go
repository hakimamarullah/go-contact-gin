package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"contact_chiv2/domain/contract"
	"contact_chiv2/domain/model"
)

type AddContactHandler struct {
	uc contract.AddContactUsecaseInterface
}

func NewAddContactHandler(usecase contract.AddContactUsecaseInterface) contract.MainHandlerInterface {
	return &AddContactHandler{
		uc: usecase,
	}
}

func (h *AddContactHandler) GetHandler() (method string, path string, handlerfuncs func(w http.ResponseWriter, r *http.Request)) {
	return http.MethodPost, "/addcontact", h.Handle
}

func (h *AddContactHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var requestBody model.AddContactRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		panic(err)
	}

	lastinserted, err := h.uc.AddContact(requestBody)

	w.Write([]byte(fmt.Sprintln("Error :", err)))
	w.Write([]byte(fmt.Sprintln("LastInserted :", lastinserted)))
}
