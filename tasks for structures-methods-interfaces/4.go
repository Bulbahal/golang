package main

import "fmt"

// 4. Встраивание структур
// Задача: Создайте структуру Person с полями Name и Age.
// Затем создайте структуру Employee, встраивающую Person, и добавляющую поле Salary.
type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	Salary int
}

func main() {
	person := Employee{Salary: 1000}
	person.Name = "James"
	person.Age = 44
	fmt.Println(person)
}
