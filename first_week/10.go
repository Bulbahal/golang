package main

import "fmt"

// Уникальные элементы
// Создайте программу, которая принимает массив и возвращает новый массив, содержащий только уникальные элементы.
func main() {
	arr := []int{1, 2, 3, 4, 5, 5}
	mapa := map[int]int{}
	result := []int{}
	for _, v := range arr {
		mapa[v]++
	}
	for k, v := range mapa {
		if v == 1 {
			result = append(result, k)

		}

	}
	fmt.Println(result)
}
