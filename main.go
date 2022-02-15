package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-rest-api/repositories"
	"github.com/go-rest-api/routers"
)

func main() {
	router := routers.GetRoutes()
	http.Handle("/", router)
	//http.ListenAndServe(":8080", router)

	product, err := repositories.GetProductByCategoryId("1")

	if err != nil {
		fmt.Println(err)
	}

	if product != nil {
		json, err := json.Marshal(product)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(json))
	}

	/* 	products, err := repositories.GetProducts()
	   	if err != nil {
	   		fmt.Println(err)
	   	}
	   	fmt.Println(products) */
}
