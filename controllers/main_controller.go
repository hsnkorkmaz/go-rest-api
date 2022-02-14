package controllers

import (
	"fmt"
	"net/http"
)

func MainController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Api main controller")
}
