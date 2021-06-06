package model

import "fmt"

type Product struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Description string   `json:"description"`
	Image       []string `json:"image"`
	Stock       int      `json:"stock"`
	Category    Category `json:"category"`
}

func (product Product) Add() Product {
	product.Id = fmt.Sprint("Pro00", index)
	index++
	// var productlist []Product
	for _, category := range categories {
		if category.Name == product.Category.Name {

			category.ProductList = append(category.ProductList, product)
			product.Category = category
		}
	}

	ProductList = append(ProductList, product)
	return product
}
func (product Product) Update(Id string) Product {
	for index, temp_product := range ProductList {
		if temp_product.Id == Id {
			product.Id = temp_product.Id
			product.Category = temp_product.Category
			ProductList[index] = product
		}
	}
	for index, temp_product := range product.Category.ProductList {
		if temp_product.Id == Id {
			product.Category.ProductList[index] = product
		}
	}
	return product
}
func (product Product) Delete(Id string) Product {
	var DeletedProduct Product
	for i, TempProduct := range ProductList {
		if TempProduct.Id == Id {
			ProductList = append(ProductList[:i], ProductList[i+1:]...)
			DeletedProduct = TempProduct
			break
		}
	}
	for index, TempProduct := range DeletedProduct.Category.ProductList {
		if TempProduct.Id == Id {
			DeletedProduct.Category.ProductList = append(DeletedProduct.Category.ProductList[:index], DeletedProduct.Category.ProductList[index+1:]...)
		}
	}
	return DeletedProduct
}
func (User Product) ViewAll() []Product {
	return ProductList
}

func (product Product) View(Id string) *Product {
	for _, temp_product := range ProductList {
		if temp_product.Id == Id {
			product = temp_product
		}
	}
	return &product
}
