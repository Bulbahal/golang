package main

import "fmt"

// 4. Создание указателя и присваивание значения
// Задача:
// Напишите функцию CreatePointer(x int) *int, которая возвращает указатель на переданное число.
//
// Пример:
//
// ptr := CreatePointer(100)
// fmt.Println(*ptr) // 100
// Дополнительно:
//
// Объясните, где в памяти будет храниться значение 100 после вызова функции
func CreatePointer(x int) *int {
	return &x
}
func main() {
	ptr := CreatePointer(100)
	fmt.Println(*ptr)
}
