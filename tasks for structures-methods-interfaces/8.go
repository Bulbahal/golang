package main

import (
	"fmt"
	"strings"
)

//8. Интерфейс для работы с БД
//Задача: Создайте интерфейс Database с методами Save(data string) и Read() string.
//Реализуйте его в структуре MockDB, которая сохраняет данные в памяти (в слайс).

type Database interface {
	Save(data string)
	Read() string
}

type MockDB struct {
	data []string
}

func (m *MockDB) Save(data string) {
	m.data = append(m.data, data)
}
func (m *MockDB) Read() string {
	return strings.Join(m.data, ", ")
}

func main() {
	r := &MockDB{}
	var database Database = r
	database.Save("Golang")
	database.Save("Python")
	database.Save("C++")
	fmt.Println(database.Read())
	fmt.Println(r.data)
}
