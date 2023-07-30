package handler

import (
	"contact_ginv1/domain/contract"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getContactByNumber struct {
	uc contract.GetContactUsecaseInterface
}

func (g *getContactByNumber) Handle(c *gin.Context) {
	contacts, err := g.uc.GetContactByNumber(c.Param("number"))
	if err != nil {
		c.Status(http.StatusNotFound)
		c.Writer.Write([]byte(fmt.Sprintf("Error %s", err.Error())))
		return
	}
	c.JSON(http.StatusOK, contacts)
	return
}

func (g *getContactByNumber) GetHandler() (method string, path string, handlerfuncs func(c *gin.Context)) {
	return http.MethodGet, "/contact/phone/:number", g.Handle
}

func NewGetContactByNumber(usecase contract.GetContactUsecaseInterface) contract.MainHandlerInterface {
	return &getContactByNumber{
		uc: usecase,
	}
}
