package main

import (
	"fmt"
)

// Фильтрация in-place
// Напишите функцию, которая удаляет из среза все четные числа in-place (без выделения нового среза).
// Пример:
// Вход: []int{1, 2, 3, 4, 5}
// Выход: [1 3 5]
func delete(slice *[]int) {
	i := 0
	for _, v := range *slice {
		if v%2 != 0 {
			(*slice)[i] = v
			i++
		}
	}
	*slice = (*slice)[:i]
}
func main() {
	slice := []int{13, 22, 31, 40}
	delete(&slice)
	fmt.Println(slice)
}
