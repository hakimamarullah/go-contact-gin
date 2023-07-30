package router

import (
	"contact_ginv1/app/provider"
	"contact_ginv1/domain/contract"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func NewServerGin() (router *gin.Engine) {
	router = gin.Default()
	var server []contract.MainHandlerInterface
	server = append(server, provider.NewGetAllHandler(), provider.NewAddContactHandler(), provider.NewGetContactByNumberHandler(), provider.NewGetContactByIMEIHandler())

	for _, handler := range server {
		method, path, handlerFunc := handler.GetHandler()
		switch method {
		case http.MethodGet:
			router.GET(path, handlerFunc)
		case http.MethodPost:
			router.POST(path, handlerFunc)
		case http.MethodPut:
			router.GET(path, handlerFunc)
		case http.MethodDelete:
			router.GET(path, handlerFunc)
		default:
			log.Fatal("unsupported request method")
		}
	}
	return
}
