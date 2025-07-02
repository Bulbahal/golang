package main

import "fmt"

// 3. Поиск дубликатов в массиве
// Дан массив чисел. Верните true, если в массиве есть дубликаты, и false в противном случае.
func main() {
	p := false
	array := []int{1, 2, 3, 4, 4}
	m := make(map[int]int)
	for _, value := range array {
		m[value]++
		if m[value] > 1 {
			p = true
		}
	}
	if p {
		fmt.Println("Дубликаты есть")
	} else {
		fmt.Println("Дубликатов нет")
	}
}
