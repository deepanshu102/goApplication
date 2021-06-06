package model

var index = 0

var userList []Users
var customerList []Customer
var categories []Category
var ProductList []Product

type ProductPurchaseStock struct {
	Id       string  `json:"id"`
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
	Amount   float64 `json:"amount"`
}

type Address struct {
	Id            string `json:"id"`
	FirstAddress  string `json:"firstAddress"`
	SecondAddress string `json:"secondAddress"`
	State         string `json:"state"`
	Country       string `json:"country"`
	Pincode       string `json:"pincode"`
}
type Shipping struct {
	Id      string  `json:"id"`
	Address Address `json:"shipping-address"`
	Status  string  `json:"status"`
}
type Orders struct {
	Id          string                 `json:"id"`
	Status      string                 `json:"status"`
	ProductList []ProductPurchaseStock `json:"product-list"`
	Amount      float64                `json:"amount"`
	Shipping    Shipping               `json:"shipping-address"`
}
