package main

import "fmt"

//  3. Обобщенное объединение:
//     Напишите функцию `Merge[T any](a, b []T) []T`, которая объединяет два среза одного типа в один новый срез.
func Merge[T any](a, b []T) []T {
	result := make([]T, 0)
	result = append(result, a...)
	result = append(result, b...)
	return result
}
func main() {
	a := []int{1, 2}
	b := []int{3, 4, 5}
	rsInt := Merge(a, b)
	fmt.Println(rsInt)
	c := []string{"hello", "world"}
	d := []string{"qq", "Earth"}
	rsString := Merge(c, d)
	fmt.Println(rsString)
}
