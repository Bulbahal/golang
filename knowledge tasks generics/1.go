package main

import "fmt"

//1. Создание обобщенной функции для поиска элемента:
//   Напишите функцию `Find[T comparable](slice []T, value T) int`, которая ищет элемент в срезе
//  и возвращает его индекс. Если элемент не найден, возвращайте -1.

func Find[T comparable](slice []T, value T) int {
	for index, item := range slice {
		if item == value {
			return index
		}
	}
	return -1
}
func main() {
	fmt.Println(Find([]string{"a", "b", "c"}, "d"))
}
