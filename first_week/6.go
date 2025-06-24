package main

import "fmt"

// 6. Поиск максимального числа
// Создайте программу, которая находит максимальное число в массиве целых чисел.

func main() {
	a := []int{1, 2, 3, 64, 4, 5}
	rs := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > rs {
			rs = a[i]
		}
	}
	fmt.Println(rs)
}
