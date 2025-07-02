package main

import "fmt"

// Найти пересечение двух массивов
// Даны два массива чисел. Верните массив их пересечения (общие элементы).
func main() {
	arr1 := []int{1, 2, 3, 4, 4, 5}
	arr2 := []int{1, 2, 3, 4, 4, 8}
	m1 := make(map[int]int)
	result := []int{}
	for _, v := range arr1 {
		m1[v]++
	}
	for _, value := range arr2 {
		if m1[value] > 0 {
			result = append(result, value)
			m1[value]--
		}
	}
	fmt.Println(result)
}
