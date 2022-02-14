package main

import (
	"net/http"

	"github.com/go-rest-api/routers"
)

func main() {
	router := routers.GetRoutes()
	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
