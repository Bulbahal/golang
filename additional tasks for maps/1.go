package main

import "fmt"

// 1. Подсчет частоты символов
// Дана строка. Посчитайте, сколько раз каждый символ встречается в строке.
func main() {
	strings := "Hello, World!"
	m := make(map[string]int)
	for _, s := range strings {
		m[string(s)]++
	}
	fmt.Println(m)
}
