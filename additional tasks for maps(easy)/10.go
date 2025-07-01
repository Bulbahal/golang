package main

import "fmt"

// Инвертирование мапы
// Дана мапа m := map[int]string{1: "one", 2: "two"}.
// Создайте новую мапу, где ключи и значения поменяются местами (map[string]int).
func main() {
	m := map[int]string{1: "one", 2: "two"}
	result := make(map[string]int)
	for k, v := range m {
		result[v] = k
	}
	fmt.Println(result)
}
