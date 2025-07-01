package main

import "fmt"

// Объединение двух срезов
// Напишите функцию, которая принимает два среза и возвращает новый срез,
// содержащий все элементы обоих без использования append (на базе make и копирования).
// Пример:
// Вход: []int{1, 2}, []int{3, 4}
// Выход: [1 2 3 4]
func copyrovanie(s1 []int, s2 []int) {
	result := make([]int, len(s1)+len(s2))
	copy(result, s1)
	copy(result[len(s1):], s2)
	fmt.Println(result)
}
func main() {
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	copyrovanie(s1, s2)
}
