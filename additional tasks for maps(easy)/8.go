package main

import "fmt"

// 8. Подсчёт частоты символов в строке
// Дана строка "hello". Создайте мапу, где ключи — символы, а значения — сколько раз они встречаются.
func main() {
	strings := "hello"
	m := make(map[string]int)
	for _, s := range strings {
		m[string(s)]++
	}
	fmt.Println(m)
}
