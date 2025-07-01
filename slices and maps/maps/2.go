package main

import "fmt"

// Объединение двух мап
// Напишите функцию, которая принимает две map[string]int и возвращает новую мапу, объединяя их. Если ключ присутствует в обеих мапах, сложите значения.
// Пример:
// Вход:
// m1 := map[string]int{"a": 1, "b": 2}
// m2 := map[string]int{"b": 3, "c": 4}
// Выход: map[a:1 b:5 c:4]
func associationMaps(m1, m2 map[string]int) map[string]int {
	m := map[string]int{}
	for k, v := range m1 {
		m[k] = v
	}
	for k, v := range m2 {
		if _, ok := m[k]; !ok {
			m[k] = v
		} else {
			m[k] = m[k] + v
		}
	}
	return m
}

func main() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	fmt.Println(associationMaps(m1, m2))
}
