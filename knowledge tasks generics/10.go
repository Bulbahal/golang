package main

import "fmt"

//10. Обобщенная функция для проверки уникальности:
//    Напишите функцию `Unique[T comparable](slice []T) []T`, которая возвращает новый срез, содержащий только уникальные элементы исходного среза.

func Unique[T comparable](slice []T) []T {
	result := make([]T, 0)
	m := make(map[T]bool)
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = true
			result = append(result, v)
		}
	}
	return result
}
func main() {
	rsInt := []int{52, 128, 64, 64, 64}
	fmt.Println(Unique(rsInt))
	rsString := []string{"banana", "meow", "apple", "mango", "meow"}
	fmt.Println(Unique(rsString))
}
