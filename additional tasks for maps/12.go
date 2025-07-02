package main

import "fmt"

//12. Найти пропущенное число
//Дан массив чисел от 1 до n с одним пропущенным числом.
//Найдите его, используя хешмапу.

func main() {
	arr := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	m := make(map[int]bool)
	n := len(arr) + 1
	for _, v := range arr {
		m[v] = true
	}
	for i := 1; i < n; i++ {
		if !m[i] {
			fmt.Println(i)
		}
	}
}
