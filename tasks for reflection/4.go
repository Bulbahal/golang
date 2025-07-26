package main

import (
	"fmt"
	"reflect"
)

// Задача 4: Итерация по полям структуры
// Создайте структуру с различными типами полей. Напишите функцию, которая итерируется по всем полям структуры с помощью рефлексии и выводит:
// - Название поля
// - Тип поля
// - Значение поля
type Car struct {
	Model string
	Age   int
}

func iteration(s interface{}) {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fmt.Println("Название поля:", field.Name)
		fmt.Println("Тип поля:", field.Type)
		fmt.Println("Значение поля:", v.Field(i))
	}
}
func main() {
	c := Car{"Audi", 2000}
	iteration(c)
}
