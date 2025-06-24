package main

import "fmt"

// 8. Сумма элементов массива
// Реализуйте программу, которая вычисляет сумму всех элементов в массиве.
func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := 0
	for i := 0; i < len(arr); i++ {
		result += arr[i]
	}
	fmt.Println(result)
}
