package main

import "fmt"

// 17. Найти разницу между двумя массивами
// Даны два массива. Верните элементы, которые есть в одном, но нет в другом.
func subsetss(arr1 []string, arr2 []string) []string {
	m := make(map[string]bool)
	result := []string{}
	for _, v := range arr1 {
		m[v] = true
	}
	for _, v := range arr2 {
		if _, ok := m[v]; ok {
			continue
		} else {
			result = append(result, v)
		}
	}
	return result
}
func main() {
	arr1 := []string{"a", "b", "c"}
	arr2 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(subsetss(arr1, arr2))
}
