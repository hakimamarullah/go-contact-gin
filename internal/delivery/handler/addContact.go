package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"contact_ginv1/domain/contract"
	"contact_ginv1/domain/model"
)

type AddContactHandler struct {
	uc contract.AddContactUsecaseInterface
}

func NewAddContactHandler(usecase contract.AddContactUsecaseInterface) contract.MainHandlerInterface {
	return &AddContactHandler{
		uc: usecase,
	}
}

func (h *AddContactHandler) GetHandler() (method string, path string, handlerfuncs func(c *gin.Context)) {
	return http.MethodPost, "/contact", h.Handle
}

func (h *AddContactHandler) Handle(c *gin.Context) {
	var requestBody model.AddContactRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		panic(err)
	}

	lastinserted, err := h.uc.AddContact(requestBody)

	w := c.Writer
	w.Write([]byte(fmt.Sprintln("Error :", err)))
	w.Write([]byte(fmt.Sprintln("LastInserted :", lastinserted)))
}
