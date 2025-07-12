package main

import "fmt"

//6. Пустой интерфейс и type assertion
//Задача: Напишите функцию PrintType, которая принимает пустой интерфейс (interface{}),
//определяет тип переданного значения и печатает его.

func PrintType(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Printf("Это int: %d\n", v)
	case string:
		fmt.Printf("Это string: %s\n", v)
	case bool:
		fmt.Printf("Это bool: %v\n", v)
	default:
		fmt.Printf("Неизвестный тип: %T\n", v)
	}
}
func main() {
	PrintType(1)
	PrintType("Hello")
	PrintType(true)
	PrintType(3.145623)
}
