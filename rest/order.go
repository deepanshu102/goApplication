package rest

import (
	"base/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	userId := vars["id"]
	ResBody, _ := ioutil.ReadAll(r.Body)
	var order model.Orders
	json.Unmarshal(ResBody, &order)
	order = model.Orders.Add(order, userId)
	json.NewEncoder(w).Encode(order)
}

func ViewOrder(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	Id := vars["id"]
	listCart := model.Orders.View(model.Orders{}, Id)
	log.Println("View ORDER LIST :- ", listCart)
	json.NewEncoder(w).Encode(listCart)
}

//Update Order refer to Cancle the order
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	userId := r.URL.Query().Get("id")
	OrderId := r.URL.Query().Get("oId")
	var Order model.Orders
	Order = model.Orders.UpdateOrder(Order, userId, OrderId)
	json.NewEncoder(w).Encode(Order)

}
