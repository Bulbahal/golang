package main

import "fmt"

//4. Обобщенная функция для подсчета элементов:
//   Создайте функцию `Count[T any](slice []T) int`, которая возвращает количество элементов в переданном срезе.

func Count[T any](slice []T) int {
	return len(slice)
}
func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []string{"aa", "bb", "cc"}
	fmt.Println(Count(s1))
	fmt.Println(Count(s2))
}
