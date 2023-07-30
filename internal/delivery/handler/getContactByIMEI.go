package handler

import (
	"contact_ginv1/domain/contract"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getContactByIMEI struct {
	uc contract.GetContactUsecaseInterface
}

func (g *getContactByIMEI) Handle(c *gin.Context) {
	contacts, err := g.uc.GetContactByIMEI(c.Param("imei"))
	if err != nil {
		c.Status(http.StatusNotFound)
		c.Writer.Write([]byte(fmt.Sprintf("Error %s", err.Error())))
		return
	}
	c.JSON(http.StatusOK, contacts)
	return
}

func (g *getContactByIMEI) GetHandler() (method string, path string, handlerfuncs func(c *gin.Context)) {
	return http.MethodGet, "/contact/imei/:imei", g.Handle
}

func NewGetContactByIMEI(usecase contract.GetContactUsecaseInterface) contract.MainHandlerInterface {
	return &getContactByIMEI{
		uc: usecase,
	}
}
