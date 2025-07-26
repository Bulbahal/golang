package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// Задача 3: Сериализация и десериализация
// Напишите функцию, которая принимает любую структуру и возвращает её сериализованный формат
// (например, JSON) с использованием рефлексии. Затем создайте другую функцию, которая принимает
// JSON-строку и десериализует её обратно в структуру.
type Book struct {
	Name  string
	Pages int
}

func ToJson(b interface{}) ([]byte, error) {
	if reflect.ValueOf(b).Kind() != reflect.Struct {
		return nil, errors.New("ожидание структуры")
	}
	return json.Marshal(b)
}
func fromJson(data []byte, b interface{}) error {
	return json.Unmarshal(data, b)
}
func main() {
	b := Book{"Martin Iden", 100}
	jsonData, err := ToJson(b)
	if err != nil {
		fmt.Println("Ошибка сериализации: " + err.Error())
	}
	fmt.Printf("JSON: %s\n", jsonData)
	var book Book
	err = fromJson(jsonData, &book)
	if err != nil {
		panic("Ошибка десериализации: " + err.Error())
	}
	fmt.Printf("Book: %s\n", book)

}
