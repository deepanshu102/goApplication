package rest

import (
	"base/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	ResBody, _ := ioutil.ReadAll(r.Body)
	var Product model.Product
	json.Unmarshal(ResBody, &Product)
	Product = model.Product.Add(Product)
	json.NewEncoder(w).Encode(Product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	Id := vars["id"]
	resBody, _ := ioutil.ReadAll(r.Body)
	var updatedProduct model.Product
	json.Unmarshal(resBody, &updatedProduct)
	updatedProduct = model.Product.Update(updatedProduct, Id)
	log.Println(updatedProduct)
	json.NewEncoder(w).Encode(updatedProduct)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	Id := vars["id"]
	var DeletedProduct model.Product
	DeletedProduct = model.Product.Delete(DeletedProduct, Id)
	json.NewEncoder(w).Encode(DeletedProduct)
}
func ViewProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	listProduct := model.Product.ViewAll(model.Product{})
	log.Println("View products :- ", listProduct)
	json.NewEncoder(w).Encode(listProduct)
}
