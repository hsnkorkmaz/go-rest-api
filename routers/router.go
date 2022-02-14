package routers

import (
	"github.com/go-rest-api/controllers"
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	router := mux.NewRouter()

	//category handlers
	router.HandleFunc("/", controllers.MainController).Methods("GET")
	router.HandleFunc("/categories", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/categories", controllers.InsertCategory).Methods("POST")
	router.HandleFunc("/categories/{id:[0-9]+}", controllers.GetCategoryById).Methods("GET")
	router.HandleFunc("/categories/{id:[0-9]+}", controllers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id:[0-9]+}", controllers.DeleteCategory).Methods("DELETE")

	return router
}
