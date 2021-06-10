package model

import "fmt"

func (User Users) Add() Users {
	User.Id = fmt.Sprint("User00", index)
	index++
	userList = append(userList, User)
	return User
}
func (User Users) Update(Id string) Users {
	for index, user := range userList {
		if user.Id == Id {
			User.Id = user.Id
			userList[index] = User
		}
	}
	return User
}
func (User Users) Delete(Id string) Users {
	var DeletedUser Users
	for i, user := range userList {
		if user.Id == Id {
			userList = append(userList[:i], userList[i+1:]...)
			DeletedUser = user
			break
		}
	}
	return DeletedUser
}
func (User Users) ViewAll() []Users {
	return userList
}

func (User Users) View(Id string) Users {
	var userDetails Users
	for _, user := range userList {
		if user.Id == Id {
			userDetails = user
		}
	}
	return userDetails
}

func (User Users) Login(email, pass string) Users {

	for _, user := range userList {
		if user.Email == email && user.Password == pass {
			User = user
			break
		}
	}
	return User

}
