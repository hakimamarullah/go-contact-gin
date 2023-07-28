package router

import (
	"contact_chiv2/app/provider"
	"contact_chiv2/domain/contract"

	"github.com/go-chi/chi/v5"
)

func NewServerChi() *chi.Mux {
	router := chi.NewRouter()

	var server []contract.MainHandlerInterface
	server = append(server, provider.NewAddContactHandler(), provider.NewGetAllHandler())

	for _, handler := range server {
		router.MethodFunc(handler.GetHandler())
	}

	return router
}
