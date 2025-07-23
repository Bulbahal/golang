package main

import "fmt"

//  6. Обобщенная функция для фильтрации:
//     Напишите функцию `Filter[T any](slice []T, predicate func(T) bool) []T`, которая возвращает новый срез,
//     содержащий только те элементы, для которых предикат возвращает true.
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if predicate(v) == true {
			result = append(result, v)
		}
	}
	return result
}
func main() {
	str := []string{"a", "bb", "ccc", "d"}
	rsStr := Filter(str, func(s string) bool {
		return len(s) > 1
	})
	fmt.Println(rsStr)
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rsInts := Filter(ints, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(rsInts)
}
