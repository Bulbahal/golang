package main

import "fmt"

//12. Конструкторы и приватные поля
//Задача:
//Создайте структуру User с приватными полями login и password.
//Напишите конструктор NewUser(login, password string) *User и метод GetLogin() string, который возвращает логин.

type User struct {
	login    string
	password string
}

func NewUser(login, password string) *User {
	return &User{login: login, password: password}
}
func (u *User) GetLogin() string {
	return u.login
}

func main() {
	user := NewUser("Marat", "qwerty123")
	fmt.Println(user.GetLogin())
}
