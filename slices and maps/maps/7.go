package main

import "fmt"

// Удаление дубликатов из среза с помощью мапы
// Напишите функцию, которая принимает срез строк и возвращает новый срез без дубликатов, используя мапу для проверки уникальности.
// Пример:
// Вход: []string{"a", "b", "a", "c"}
// Выход: []string{"a", "b", "c"}
func uniqueness(s []string) []string {
	seems := make(map[string]bool)
	result := []string{}
	for _, v := range s {
		if !seems[v] {
			seems[v] = true
			result = append(result, v)
		}
	}
	return result
}
func main() {
	s := []string{"a", "b", "a", "c", "c", "c", "g"}
	fmt.Println(uniqueness(s))
}
