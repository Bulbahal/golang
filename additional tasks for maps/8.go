package main

import "fmt"

//Найти самый частый элемент
//Дан массив чисел. Найдите число, которое встречается чаще всего.

func main() {
	arr := []int{1, 2, 3, 4, 4, 4, 4, 4, 3, 5, 5, 5, 5, 5, 5}
	m := map[int]int{}
	for _, v := range arr {
		m[v]++
	}
	max := 0
	for _, k := range m {
		if m[k] > max {
			max = k
		}
	}
	fmt.Println(max)
}
