package main

import "fmt"

//7. Подсчет суммы двух чисел
//Дан массив чисел и целевое число.
//Проверьте, есть ли в массиве два числа, сумма которых равна целевому.

func main() {
	arr1 := []int{6, 2, 2, 4, 5}
	n := 8
	m := make(map[int]bool)
	result := false
	for _, v := range arr1 {
		if m[n-v] {
			result = true
		}
		m[v] = true
	}
	fmt.Println(result)
}
