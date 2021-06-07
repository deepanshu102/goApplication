package main

import (
	"base/rest"
)

func main() {
	rest.Handler()

	// dataInsertion()
}

// func dataInsertion() {
// 	userFile, _ := ioutil.ReadFile("data.json")
// 	// customerData, _ := ioutil.ReadFile("cartData.json")
// 	var userList []model.Users
// 	var customer model.Customer
// 	fmt.Printf("%v", customer)
// 	json.Unmarshal([]byte(userFile), &userList)
// 	// json.Unmarshal([]byte(customerData), &customer)
// 	for i := 0; i < len(userList); i++ {
// 		fmt.Printf("%+v\n", userList[i])

// 	}
// 	// fmt.Printf("%+v", customer)

// }
