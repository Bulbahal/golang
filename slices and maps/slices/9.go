package main

import (
	"fmt"
)

// Вставка элемента по индексу
// Напишите функцию, которая вставляет элемент в срез по указанному индексу с сохранением исходного capacity.
// Пример:
// Вход: slice = []int{10, 20, 30}, index = 1, value = 99
// Выход: [10 99 20 30]
func replace(value int, index int) {
	slice := []int{10, 20, 30}
	result := make([]int, len(slice))
	copy(result, slice)
	slice = append(slice[:index], value)
	slice = append(slice, result[index:]...)
	fmt.Println(slice)

}
func main() {
	value := 99
	index := 1
	replace(value, index)
}
