package model

import "fmt"

func (product Product) Add() Product {
	product.Id = fmt.Sprint("Pro00", index)
	index++
	// var productlist []Product

	for index, category := range categories {
		if category.Name == product.Category.Name {

			category.ProductList = append(category.ProductList, product)
			product.Category = category
			categories[index] = category
		}
	}

	ProductList = append(ProductList, product)
	return product
}
func (product Product) Update(Id string) Product {
	for index, temp_product := range ProductList {
		if temp_product.Id == Id {
			product.Id = temp_product.Id
			if product.Category.Name == temp_product.Category.Name {
				product.Category = temp_product.Category
				ProductList[index] = product
			} else {
				fmt.Println("Category try to change of ", product)
			}
		}
	}
	for _, category := range categories {
		if category.Name == product.Category.Name {
			for productIndex, TempProduct := range category.ProductList {
				if TempProduct.Id == product.Id {
					category.ProductList[productIndex] = product
				}
			}
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
	for index, category := range categories {
		if category.Name == DeletedProduct.Category.Name {
			for productIndex, TempProduct := range category.ProductList {
				if TempProduct.Id == DeletedProduct.Id {
					category.ProductList = append(category.ProductList[:productIndex], category.ProductList[productIndex+1:]...)
				}
			}
			categories[index] = category
		}
	}
	return DeletedProduct
}
func (User Product) ViewAll() []Product {
	return ProductList
}

func (product Product) View(Id string) Product {
	for _, temp_product := range ProductList {
		if temp_product.Id == Id {
			product = temp_product
		}
	}
	return product
}
