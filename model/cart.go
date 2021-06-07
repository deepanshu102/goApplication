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
			User.Cart.ProductList = append(user.Cart.ProductList, productPurchaseStock)
			User.Cart.Bill += productPurchaseStock.Amount
			userList[i] = User //golbal update
		}
	}

	return User.Cart

}

func (cart Cart) Delete(UserId string, Id string) Cart {
	var user Users
	user = Users.View(user, UserId)
	for index, product := range user.Cart.ProductList {
		if product.Id == Id {
			user.Cart.ProductList = append(user.Cart.ProductList[:index], user.Cart.ProductList[index+1:]...)
			user.Cart.Bill = user.Cart.Bill - product.Amount
		}
	}
	return user.Cart

}
func (cart Cart) Update(UserId string, productPurchaseStock ProductPurchaseStock) Cart {
	var cardProductList []ProductPurchaseStock
	for userIndex, users := range userList {
		if users.Id == UserId {

			cart = users.Cart
			cardProductList = cart.ProductList
			for index, cartproduct := range cardProductList {
				if productPurchaseStock.Id == cartproduct.Id {
					product := cartproduct.Product

					productPurchaseStock.Product = cartproduct.Product
					fmt.Printf("Product:---\n%+v\n", productPurchaseStock.Product)
					productPurchaseStock.Amount = product.Price * float64(productPurchaseStock.Quantity)

				}
				cardProductList[index] = productPurchaseStock
				cart.ProductList = cardProductList
				cart.Bill = cart.Bill + productPurchaseStock.Amount - cartproduct.Amount

			}
			userList[userIndex].Cart = cart
			break
		}
	}

	return cart
}

func (cart Cart) View(UserId string) Cart {
	fmt.Println("Cart View Called", UserId)
	var user Users
	user = Users.View(user, UserId)
	return user.Cart
}
