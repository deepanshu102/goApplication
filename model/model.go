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
	ProductList []ProductPurchaseStock `json:"productList"`
	Amount      float64                `json:"amount"`
	Shipping    Shipping               `json:"shippingAddress"`
}

type Product struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Description string   `json:"description"`
	Image       []string `json:"image"`
	Stock       int      `json:"stock"`
	Category    Category `json:"category"`
}
type Cart struct {
	Id          string                 `json:"id"`
	ProductList []ProductPurchaseStock `json:"productList"`
	Bill        float64                `json:"bill"`
}

type Users struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Role      string   `json:"role"`
	Cart      Cart     `json:"cart"`
	OrderList []Orders `json:"orderList"`
}

type Customer struct {
	Id string `json:"Id"`
}
type Category struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	ProductList []Product `json:"productList"`
}
