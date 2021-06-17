package rest

import (
	"base/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*

	View All the Category
*/

func CategoryList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	CategoryList := model.Category.ViewAll(model.Category{})
	log.Printf("Category List:-%+v", CategoryList)
	json.NewEncoder(w).Encode(CategoryList)
}

/*
	Create new Category
*/
func createCategory(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	ResBody, _ := ioutil.ReadAll(r.Body)
	var category model.Category
	json.Unmarshal(ResBody, &category)
	category = model.Category.Add(category)
	log.Printf("Category Created:-%+v", category)
	json.NewEncoder(w).Encode(category)
	// fmt.FPrintf(w, categorys)

}

/*
	Deleted the Category
	passing
	- Unique Id with the URL
*/

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	Id := vars["id"]
	var deletedCategory model.Category
	deletedCategory = model.Category.Delete(deletedCategory, Id)
	log.Printf("delete Category:-%+v", deletedCategory)
	json.NewEncoder(w).Encode(deletedCategory)
}

/*
	Updated Category
	1- Unique Id with URL
	2- sending new Category Body in json
*/
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	id := vars["id"]
	ResBody, _ := ioutil.ReadAll(r.Body)
	var updateCategory model.Category
	json.Unmarshal(ResBody, &updateCategory)
	updateCategory = model.Category.Update(updateCategory, id)
	json.NewEncoder(w).Encode(updateCategory)

}
