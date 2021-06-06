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

// func ProductHandler() {
// 	log.Println("Product Handler added")
// 	myRouter.HandleFunc("/product", ViewProduct).Methods("GET")
// 	myRouter.HandleFunc("/product", CreateProduct).Methods("POST")
// 	myRouter.HandleFunc("/product", UpdateProduct).Methods("PUT")
// 	myRouter.HandleFunc("/product", DeleteProduct).Methods("DELETE")

// }

func UserHandler() {
	log.Println("User Handler added")
	myRouter.HandleFunc("/userall", ViewAllUser).Methods("GET")
	myRouter.HandleFunc("/user/{id}", SingleUser).Methods("GET")
	myRouter.HandleFunc("/user", CreateUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

}

func Handler() {
	CategoryHandler()
	// ProductHandler()
	UserHandler()
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
