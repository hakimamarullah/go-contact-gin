package contract

import "net/http"

type MainHandlerInterface interface {
	Handle(w http.ResponseWriter, r *http.Request)
	GetHandler() (method string, path string, handlerfuncs func(w http.ResponseWriter, r *http.Request))
}
