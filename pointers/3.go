package main

import "fmt"

// 3. Проверка на nil
// Задача:
// Напишите функцию IsNil(val *int) bool, которая проверяет, является ли переданный указатель nil.
//
// Пример:
//
// var ptr *int
// fmt.Println(IsNil(ptr)) // true
//
// num := 42
// fmt.Println(IsNil(&num)) // false
// Подсказка:
//
// В Go nil — это нулевое значение для указателей.

func IsNil(val *int) bool {
	if val == nil {
		return true
	} else {
		return false
	}
}
func main() {
	var ptr *int
	fmt.Println(IsNil(ptr))
	num := 42
	fmt.Println(IsNil(&num))
}
