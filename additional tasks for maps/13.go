package main

import "fmt"

// 13. Проверить, является ли массив подмножеством другого
// Даны два массива. Проверьте, все ли элементы первого массива присутствуют во втором.
func subsets(arr1 []string, arr2 []string) bool {
	m := make(map[string]bool)
	for _, v := range arr2 {
		m[v] = true
	}
	for _, v := range arr1 {
		if !m[v] {
			return false
		}
	}
	return true
}
func main() {
	arr1 := []string{"a", "b", "g"}
	arr2 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(subsets(arr1, arr2))
}
