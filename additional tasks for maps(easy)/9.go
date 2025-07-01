package main

import "fmt"

// Поиск дубликатов в слайсе
// Дан слайс nums := []int{1, 2, 3, 2, 4, 5, 4}.
// Проверьте, есть ли в нём дубликаты, используя мапу.
func main() {
	dublicates := false
	nums := []int{1, 2, 3, 2, 4, 5, 4}
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
		if m[num] > 1 {
			dublicates = true
			break
		}
	}
	if dublicates {
		fmt.Println("Дубликаты есть")
	} else {
		fmt.Println("Дубликтов нет")
	}
}
