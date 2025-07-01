package main

import "fmt"

// Объединение двух мап
// Даны две мапы:
//
// m1 := map[string]int{"a": 1, "b": 2}
// m2 := map[string]int{"b": 3, "c": 4}
// Создайте новую мапу, объединив их (при конфликте берётся значение из m2).
func main() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	result := make(map[string]int)
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		result[k] = v
	}
	fmt.Println(result)
}
