package main

import "fmt"

//8. Обобщенная функция для преобразования:
//   Создайте функцию `Map[T any, U any](slice []T, transform func(T) U) []U`,
//  которая применяет функцию преобразования к каждому элементу и возвращает новый срез.

func main() {
	numbers := []int{1, 2, 3, 4}
	strings := mapa(numbers, func(n int) string {
		return fmt.Sprintf("number-%d", n)
	})
	fmt.Println(strings)
}
func mapa[T any, U any](slice []T, transform func(T) U) []U {
	result := make([]U, len(slice))
	for i := 0; i < len(slice); i++ {
		result[i] = transform(slice[i])
	}
	return result
}
