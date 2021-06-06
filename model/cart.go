package model

import "fmt"

type Cart struct {
	Id          string                 `json:"id"`
	ProductList []ProductPurchaseStock `json:"productList"`
	Bill        float64                `json:"bill"`
}

func (cart Cart) Add(UserId string, ProductId string, quantity int) Cart {
	var product *Product
	product = Product.View(*product, ProductId)

	var productPurchaseStock ProductPurchaseStock
	productPurchaseStock.Id = fmt.Sprint("PPS", index)
	index++
	productPurchaseStock.Product = *product
	productPurchaseStock.Quantity = quantity
	productPurchaseStock.Amount = product.Price * float64(quantity)

	var User *Users
	User = Users.View(*User, UserId)
	User.customer.cart.ProductList = append(cart.ProductList, productPurchaseStock)
	User.customer.cart.Bill += productPurchaseStock.Amount
	return User.customer.cart

}

func (cart Cart) Delete(UserId string, Id string) {
	var user *Users
	user = Users.View(*user, UserId)
	for index, product := range user.customer.cart.ProductList {
		if product.Id == Id {
			user.customer.cart.ProductList = append(user.customer.cart.ProductList[:index], user.customer.cart.ProductList[index+1:]...)
			user.customer.cart.Bill = user.customer.cart.Bill - product.Amount
		}
	}

}
func (cart Cart) update(UserId string, productPurchaseStock ProductPurchaseStock) {
	var user *Users
	user = Users.View(*user, UserId)
	for index, product := range user.customer.cart.ProductList {
		if product.Id == productPurchaseStock.Id {
			user.customer.cart.Bill -= product.Amount
			productPurchaseStock.Amount = productPurchaseStock.Product.Price * float64(productPurchaseStock.Quantity)
		}
	}
}
