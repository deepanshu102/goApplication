package rest

import (
	"base/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ViewAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	userList := model.Users.ViewAll(model.Users{})
	log.Printf("%+v", userList)
	json.NewEncoder(w).Encode(userList)

}
func SingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	Id := vars["id"]
	json.NewEncoder(w).Encode(model.Users.View(model.Users{}, Id))

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	resBody, _ := ioutil.ReadAll(r.Body)
	var user model.Users
	json.Unmarshal(resBody, &user)
	user = model.Users.Add(user)
	log.Printf("Created:-\n%+v", user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var User model.Users
	json.Unmarshal(reqBody, &User)
	User = model.Users.Update(User, id)
	log.Printf("Updated:-\n%+v", User)
	json.NewEncoder(w).Encode(User)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var user model.Users
	user = model.Users.Delete(user, id)
	log.Printf("Deleted User Details :- %+v", user)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("pass")
	var User model.Users
	User = model.Users.Login(User, email, password)
	log.Printf("Login By :-%+v", User)
	json.NewEncoder(w).Encode(User)
}
