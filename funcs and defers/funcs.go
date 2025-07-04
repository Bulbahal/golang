package main

import (
	"fmt"
)

// 1. Подсчёт частоты элементов в слайсе
// Напишите функцию CountFrequency, которая принимает слайс строк и возвращает мапу,
// где ключи — элементы слайса, а значения — количество их вхождений.
//func CountFrequency(slice []string) map[string]int {
//	count := make(map[string]int)
//	for _, v := range slice {
//		count[v]++
//	}
//	return count
//}
//func main() {
//	str := []string{"a", "b", "c"}
//	fmt.Println(CountFrequency(str))
//}

// 2. Фильтрация слайса по условию
// Напишите функцию FilterEven, которая принимает слайс целых чисел и возвращает новый слайс, содержащий только чётные числа.
//func FilterEven(nums []int) []int {
//	result := []int{}
//	for n := range nums {
//		if nums[n]%2 == 0 {
//			result = append(result, nums[n])
//		}
//	}
//	return result
//}
//func main() {
//	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
//	fmt.Println(FilterEven(nums))
//}

// 3. Объединение двух мап
// Напишите функцию MergeMaps, которая принимает две мапы (map[string]int) и возвращает новую мапу, объединяющую их.
// Если ключ присутствует в обеих мапах, выбирается значение из второй.
//func MergeMaps(m1, m2 map[string]int) map[string]int {
//	result := make(map[string]int)
//	for k, v := range m1 {
//		result[k] = v
//	}
//	for k, v := range m2 {
//		result[k] = v
//	}
//	return result
//}
//func main() {
//	m1 := map[string]int{"a": 1, "b": 1}
//	m2 := map[string]int{"b": 2, "c": 3}
//	fmt.Println(MergeMaps(m1, m2))
//}

// 4. Удаление дубликатов из слайса
// Напишите функцию RemoveDuplicates, которая принимает слайс строк и возвращает новый слайс без дубликатов.
//func RemoveDuplicates(str []string) []string {
//	result := make([]string, 0)
//	m := make(map[string]int)
//	for _, v := range str {
//		m[v]++
//		if m[v] == 1 {
//			result = append(result, v)
//		}
//	}
//	return result
//}
//func main() {
//	str := []string{"a", "b", "c", "c"}
//	fmt.Println(RemoveDuplicates(str))
//}

// 5. Напишите функцию InvertMap, которая принимает мапу map[string]int и возвращает новую мапу map[int]string,
// где ключи и значения поменялись местами. Если в исходной мапе есть дублирующиеся значения, вернуть ошибку.
//func InvertMap(m map[string]int) (map[int]string, error) {
//	result := make(map[int]string)
//	for k, v := range m {
//		result[v] = k
//	}
//	if len(result) != len(m) {
//		return nil, errors.New("Есть дубликаты")
//	} else {
//		return result, nil
//	}
//}
//func main() {
//	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
//	m1 := map[string]int{"one": 1, "two": 2, "three": 1, "four": 4}
//	fmt.Println(InvertMap(m))
//	fmt.Println(InvertMap(m1))
//}

// 6. Разделение слайса на пакеты (чанки)
// Напишите функцию ChunkSlice, которая принимает слайс целых чисел и размер чанка
// а возвращает слайс слайсов, разбитый на части указанного размера.
//func ChunkSlice(slice []int, chunkSize int) [][]int {
//	result := [][]int{}
//	if len(slice) < chunkSize {
//		return append(result, slice)
//	}
//	for i := 0; i < len(slice); i += chunkSize {
//		end := i + chunkSize
//		if end > len(slice) {
//			end = len(slice)
//		}
//		result = append(result, slice[i:end])
//	}
//	return result
//
//}
//func main() {
//	slice := []int{6, 2, 2, 4, 5, 3}
//	chunks := 4
//	fmt.Println(ChunkSlice(slice, chunks))
//}

// 7. Поиск уникальных элементов в слайсе
// Напишите функцию FindUnique, которая принимает слайс целых чисел и возвращает новый слайс,
// содержащий только уникальные элементы (встречающиеся 1 раз).
//func FindUnique(nums []int) []int {
//	m := make(map[int]int)
//	result := make([]int, 0)
//	for _, v := range nums {
//		m[v]++
//		if m[v] == 1 {
//			result = append(result, v)
//		}
//	}
//	return result
//}
//func main() {
//	nums := []int{1, 2, 3, 4, 4, 5}
//	fmt.Println(FindUnique(nums))
//}

// 8. Подсчёт среднего значения в слайсе
// Напишите функцию Average, которая принимает слайс чисел и возвращает их среднее значение.
// Если слайс пуст, возвращает 0.
//func Average(x []int) float64 {
//	result := float64(0)
//	for _, v := range x {
//		result += float64(v)
//	}
//	return result / float64(len(x))
//}
//func main() {
//	nums := []int{1, 2, 3, 4, 5}
//	fmt.Println(Average(nums))
//}

// 9. Группировка строк по длине
// Напишите функцию GroupByLength, которая принимает слайс строк и возвращает мапу, где ключ — длина строки,
// а значение — слайс строк этой длины.
func GroupByLength(str []string) map[int][]string {
	m := make(map[int][]string)
	for _, v := range str {
		m[len(v)] = append(m[len(v)], v)

	}
	return m
}
func main() {
	str := []string{"hello", "world", "goodbye"}
	fmt.Println(GroupByLength(str))
}
