package main

import "fmt"

// 18. Подсчет количества пар с заданной суммой
// Дан массив чисел и число k. Посчитайте количество пар чисел, сумма которых равна k.
func main() {
	arr1 := []int{6, 2, 2, 4, 5, 3}
	k := 8
	m := make(map[int]bool)
	//result := false
	index := 0
	for _, v := range arr1 {
		if m[k-v] {
			index++
			//result = true
		}
		m[v] = true
	}
	//fmt.Println(result)
	fmt.Println(index)
}
