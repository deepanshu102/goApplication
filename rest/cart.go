package rest

import (
	"base/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateCart(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	userId := vars["id"]
	ResBody, _ := ioutil.ReadAll(r.Body)
	var cart model.Cart
	json.Unmarshal(ResBody, &cart)
	cart = model.Cart.Add(cart, userId)
	json.NewEncoder(w).Encode(cart)
}

func UpdateCart(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	Id := vars["id"]
	resBody, _ := ioutil.ReadAll(r.Body)
	var updatedCart model.Cart
	var productPurchaseStock model.ProductPurchaseStock
	json.Unmarshal(resBody, &productPurchaseStock)
	updatedCart = model.Cart.Update(updatedCart, Id, productPurchaseStock)
	log.Println(updatedCart)
	json.NewEncoder(w).Encode(updatedCart)
}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	UserId := r.URL.Query().Get("Id")
	ProductId := r.URL.Query().Get("pro")
	var DeletedCart model.Cart
	DeletedCart = model.Cart.Delete(DeletedCart, UserId, ProductId)
	json.NewEncoder(w).Encode(DeletedCart)
}
func ViewCarts(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	Id := vars["id"]
	listCart := model.Cart.View(model.Cart{}, Id)
	log.Println("View Carts :- ", listCart)
	json.NewEncoder(w).Encode(listCart)
}
