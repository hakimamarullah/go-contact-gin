package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"contact_ginv1/domain/contract"
)

type getAllContactHandler struct {
	uc contract.GetAllContactUsecaseInterface
}

// GetHandler implements contract.MainHandlerInterface.
func (h *getAllContactHandler) GetHandler() (method string, path string, handlerfuncs func(c *gin.Context)) {
	return http.MethodGet, "/contact", h.Handle
}

// Handle implements contract.MainHandlerInterface.
func (h *getAllContactHandler) Handle(c *gin.Context) {
	response, err := h.uc.GetAllContact()
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
	}

	c.JSON(http.StatusOK, response)
	return
}

func NewGetAllContactHandler(usecase contract.GetAllContactUsecaseInterface) contract.MainHandlerInterface {
	return &getAllContactHandler{
		uc: usecase,
	}
}
