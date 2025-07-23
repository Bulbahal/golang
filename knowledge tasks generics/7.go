package main

import (
	"fmt"
)

//  7. Обобщенная функция для сортировки:
//     Реализуйте функцию `Sort[T comparable](slice []T)`, которая сортирует срез элементов любого типа, поддерживающего сравнение.
func Sort[T comparable](slice []T) []T {
	n := len(slice)
	result := make([]T, n)
	copy(result, slice)
	switch any(result).(type) {
	case []int:
		for i := 0; i < n-1; i++ {
			for j := 0; j < n-i-1; j++ {
				if any(result[j]).(int) > any(result[j+1]).(int) {
					result[j], result[j+1] = result[j+1], result[j]
				}
			}
		}

	case []string:
		for i := 0; i < n-1; i++ {
			for j := 0; j < n-i-1; j++ {
				if any(result[j]).(string) > any(result[j+1]).(string) {
					result[j], result[j+1] = result[j+1], result[j]
				}
			}
		}
	default:
		panic(fmt.Sprintf("unsupported type: %T", result))
	}
	return result
}
func main() {
	rs1 := []int{52, 128, 64}
	fmt.Println(Sort(rs1))
	rs2 := []string{"banana", "meow", "apple"}
	fmt.Println(Sort(rs2))
}
