package model

import "fmt"

func (order Orders) Add(UserId string) Orders {
	order.Id = fmt.Sprint("order", index)
	index++
	fmt.Print("ORDER DETAILS", order)
	var User Users
	for i, user := range userList {
		if user.Id == UserId {

			User = userList[i]
			order.ProductList = User.Cart.ProductList
			order.Amount = User.Cart.Bill
			order.Status = "Deliver Soon"
			User.OrderList = append(User.OrderList, order)
			User.Cart = Cart{}
			userList[i] = User //golbal update
			break
		}
	}

	return order

}

// func (cart Cart) Delete(UserId string, Id string) Cart {
// 	fmt.Print("Delete functionCalled")
// 	var user Users
// 	for userIndex, users := range userList {

// 		if users.Id == UserId {
// 			user = users
// 			productList := user.Cart.ProductList

// 			for index, product := range productList {
// 				if product.Id == Id {
// 					fmt.Printf("\n\n\n Found :%+v\n\n%+v\n\n", product, user)
// 					productList = append(productList[:index], productList[index+1:]...)
// 					user.Cart.Bill = user.Cart.Bill - product.Amount
// 				}
// 			}
// 			user.Cart.ProductList = productList
// 			userList[userIndex] = user
// 		}
// 	}
// 	return user.Cart

// }

func (orders Orders) UpdateOrder(UserId, orderId string) Orders {
	for index, user := range userList {
		if user.Id == UserId {
			for Oindex, order := range user.OrderList {
				if order.Id == orderId {
					userList[index].OrderList[Oindex].Status = "cancel"
					orders = userList[index].OrderList[Oindex]
					break
				}
			}
			break
		}

	}
	return orders
}

func (order Orders) View(UserId string) []Orders {
	var user Users
	user = Users.View(user, UserId)
	return user.OrderList
}
