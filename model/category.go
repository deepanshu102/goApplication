package model

import "fmt"

type Category struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	ProductList []Product `json:"productList"`
}

func (category Category) Add() Category {
	category.Id = fmt.Sprint("Cat00", index)

	index++
	categories = append(categories, category)
	return category
}
func (category Category) Update(Id string) Category {
	for index, tempCategory := range categories {
		if tempCategory.Id == Id {
			category.Id = tempCategory.Id
			categories[index] = category
		}
	}
	return category
}
func (category Category) Delete(Id string) Category {
	var DeletedUser Category
	for i, tempCategory := range categories {
		if tempCategory.Id == Id {
			categories = append(categories[:i], categories[i+1:]...)
			DeletedUser = tempCategory
			break
		}
	}
	return DeletedUser
}
func (Category Category) ViewAll() []Category {
	return categories
}

func (User Category) View(Id string) Category {
	var categoryDetails Category
	for _, tempCategory := range categories {
		if tempCategory.Id == Id {
			categoryDetails = tempCategory
		}
	}
	return categoryDetails
}
