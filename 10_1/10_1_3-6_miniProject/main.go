package main

import (
	"fmt"
)

type UsersStorage interface {
	GetAll() []User
	TryGetById(id int) (*User, bool)
	Add(user User)
	UpdateById(id int, name string, age int) bool
	RemoveById(id int) bool
}

type User struct {
	Id   int
	Name string
	Age  int
}

type UsersInMemoryStorage struct {
	users []User
}

func (u UsersInMemoryStorage) GetAll() []User {
	return u.users
}

func (u *UsersInMemoryStorage) TryGetById(id int) (*User, bool) {
	for _, user := range u.users {
		if user.Id == id {
			return &user, true
		}
	}
	return nil, false
}

func (u *UsersInMemoryStorage) Add(users User) {
	u.users = append(u.users, users) //здесь
}

func (u *UsersInMemoryStorage) UpdateById(id int, name string, age int) bool {
	for i, user := range u.users {
		if user.Id == id {
			u.users[i].Name = name
			u.users[i].Age = age
			return true
		}
	}
	return false
}

func (u *UsersInMemoryStorage) RemoveById(id int) bool {
	for i, user := range u.users {
		if user.Id == id {
			u.users = append(u.users[:i], u.users[i+1:]...)
			return true
		}
	}
	return false
}

func main() {
	var storage UsersStorage = &UsersInMemoryStorage{} //и здесь

	users := storage.GetAll()
	fmt.Println(len(users) == 0)

	user, found := storage.TryGetById(5)
	fmt.Println(!found)

	storage.Add(User{Id: 1, Name: "Josef", Age: 26})
	storage.Add(User{Id: 2, Name: "Mark", Age: 27})
	users = storage.GetAll()
	fmt.Println(len(users) == 2)

	user, found = storage.TryGetById(2)
	fmt.Println(found && user.Name == "Mark")

	fmt.Println(storage.UpdateById(2, "Andrew", 27))
	fmt.Println(storage.RemoveById(3))
	fmt.Println(storage.RemoveById(2))
}
