package main

import (
	"net/http"

	"contact_chiv2/app/router"
)

func main() {
	apps := router.NewServerChi()
	http.ListenAndServe(":8080", apps)
}
