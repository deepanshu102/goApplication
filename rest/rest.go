package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var myRouter = mux.NewRouter().StrictSlash(true)

/*Category adding for the Rest Api*/
func CategoryHandler() {
	log.Print("category Rest Handler Added")
	myRouter.HandleFunc("/category", CategoryList).Methods("GET")
	myRouter.HandleFunc("/category", createCategory).Methods("POST")
	myRouter.HandleFunc("/category/{id}", DeleteCategory).Methods("DELETE")
	myRouter.HandleFunc("/category/{id}", UpdateCategory).Methods("PUT")
}

func ProductHandler() {
	log.Println("Product Handler added")
	myRouter.HandleFunc("/product", ViewProducts).Methods("GET")
	myRouter.HandleFunc("/product", CreateProduct).Methods("POST")
	myRouter.HandleFunc("/product/{id}", UpdateProduct).Methods("PUT")
	myRouter.HandleFunc("/product/{id}", DeleteProduct).Methods("DELETE")

}

func UserHandler() {
	log.Println("User Handler added")
	myRouter.HandleFunc("/userall", ViewAllUser).Methods("GET")
	myRouter.HandleFunc("/user/{id}", SingleUser).Methods("GET")
	myRouter.HandleFunc("/user", CreateUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

}
func CartHandler() {
	log.Println("cart Handler added")
	myRouter.HandleFunc("/cart/{id}", ViewCarts).Methods("GET")
	myRouter.HandleFunc("/cart/{id}", CreateCart).Methods("POST")
	myRouter.HandleFunc("/cart/{id}", UpdateCart).Methods("PUT")
	myRouter.HandleFunc("/cart/{id}", DeleteCart).Methods("DELETE")

}
func OrderHandler() {
	log.Println("Order Handler added")
	myRouter.HandleFunc("/order/{id}", ViewOrder).Methods("GET")
	myRouter.HandleFunc("/order/{id}", CreateOrder).Methods("POST")
}

func Handler() {

	CategoryHandler()
	ProductHandler()
	UserHandler()
	CartHandler()
	OrderHandler()
	log.Println(http.ListenAndServe(":8080", myRouter))
}
