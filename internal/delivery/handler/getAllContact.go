package handler

import (
	"encoding/json"
	"net/http"

	"contact_chiv2/domain/contract"
)

type getAllContactHandler struct {
	uc contract.GetAllContactUsecaseInterface
}

// GetHandler implements contract.MainHandlerInterface.
func (h *getAllContactHandler) GetHandler() (method string, path string, handlerfuncs func(w http.ResponseWriter, r *http.Request)) {
	return http.MethodGet, "/contact", h.Handle
}

// Handle implements contract.MainHandlerInterface.
func (h *getAllContactHandler) Handle(w http.ResponseWriter, r *http.Request) {
	response, err := h.uc.GetAllContact()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func NewGetAllContactHandler(usecase contract.GetAllContactUsecaseInterface) contract.MainHandlerInterface {
	return &getAllContactHandler{
		uc: usecase,
	}
}
