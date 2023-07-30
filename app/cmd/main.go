package main

import (
	"log"
	"net/http"

	"contact_ginv1/app/router"
)

func main() {
	apps := router.NewServerGin()
	err := http.ListenAndServe("localhost:8080", apps)
	if err != nil {
		log.Fatal(err.Error())
	}
}
