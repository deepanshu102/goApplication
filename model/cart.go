package model

import "fmt"

func (cart Cart) Add(UserId string) Cart {

	var product Product
	product = Product.View(product, cart.ProductList[0].Product.Id)

	fmt.Printf("Add Cart()->%+v\t%+v", cart, product)
	var productPurchaseStock ProductPurchaseStock
	productPurchaseStock.Id = fmt.Sprint("PPS", index)
	index++
	productPurchaseStock.Product = product
	productPurchaseStock.Quantity = cart.ProductList[0].Quantity
	productPurchaseStock.Amount = product.Price * float64(cart.ProductList[0].Quantity)

	var User Users
	for i, user := range userList {
		if user.Id == UserId {
			User = userList[i]
			User.customer.cart.ProductList = append(user.customer.cart.ProductList, productPurchaseStock)
			User.customer.cart.Bill += productPurchaseStock.Amount
			userList[i] = User //golbal update
		}
	}

	return User.customer.cart

}

func (cart Cart) Delete(UserId string, Id string) Cart {
	var user Users
	user = Users.View(user, UserId)
	for index, product := range user.customer.cart.ProductList {
		if product.Id == Id {
			user.customer.cart.ProductList = append(user.customer.cart.ProductList[:index], user.customer.cart.ProductList[index+1:]...)
			user.customer.cart.Bill = user.customer.cart.Bill - product.Amount
		}
	}
	return user.customer.cart

}
func (cart Cart) Update(UserId string, productPurchaseStock ProductPurchaseStock) Cart {
	fmt.Println("Update method Called")
	fmt.Printf("%+v\n%+v\n", UserId, productPurchaseStock)
	var user Users
	var cartRef Cart
	for index, users := range userList {
		if users.Id == UserId {
			user = userList[index]
			cartRef = user.customer.cart
			for indexer, product := range cartRef.ProductList {
				if product.Id == productPurchaseStock.Id {
					productPurchaseStock.Id = product.Id
					cartRef.Bill -= product.Amount
					productPurchaseStock.Amount = productPurchaseStock.Product.Price * float64(productPurchaseStock.Quantity)
					cartRef.ProductList[indexer] = productPurchaseStock
					cartRef.Bill += productPurchaseStock.Amount
				}
			}
			userList[index] = user
		}
	}
	fmt.Println(user.customer.cart)
	return user.customer.cart
}

func (cart Cart) View(UserId string) Cart {
	fmt.Println("Cart View Called", UserId)
	var user Users
	user = Users.View(user, UserId)
	return user.customer.cart
}
