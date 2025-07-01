package main

import "fmt"

// Изменение in-place
// Напишите функцию, которая принимает срез целых чисел и умножает каждый его элемент на 2 без возврата нового среза (изменя исходный).
// Пример:
// Вход: []int{1, 2, 3}
// Выход: [2 4 6]
func cucumber(slices []int) {
	for i, v := range slices {
		(slices)[i] = v * 2
	}
}
func main() {
	slices := []int{1, 2, 3}
	cucumber(slices)
	fmt.Println(slices)
}
