package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Fprintf(w, "Return all categories")
}

func GetCategoryById(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Return category by id:"+id)

	//category := getCategory(id)
	//json.NewEncoder(w).Encode(article)
}

func InsertCategory(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var newCategory Category
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newCategory)
	if err != nil {
		fmt.Fprintf(w, "insert error:"+err.Error())
		return
	}

	fmt.Fprintf(w, "insert category type:"+r.Method)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "update category id: "+id)

	fmt.Fprintf(w, "type:"+r.Method)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "delete category id: "+id)
}
