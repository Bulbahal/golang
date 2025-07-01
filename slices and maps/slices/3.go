package main

import "fmt"

// Удаление элемента по индексу
// Напишите функцию, которая удаляет элемент из среза по заданному индексу in-place (без возврата нового среза).
// Пример:
// Вход: slice = []int{10, 20, 30, 40}, index = 1
// Выход: [10 30 40]
func d(slice *[]int, index int) {
	for i := index; i < len(*slice)-1; i++ {
		(*slice)[i] = (*slice)[i+1]
	}
	*slice = (*slice)[:len((*slice))-1]
}
func main() {
	slice := []int{10, 20, 30, 40}
	index := 1
	d(&slice, index)
	fmt.Println(slice)
}
