package main

import "fmt"

// Преобразование мапы в срез ключей/значений
// Напишите функцию, которая принимает map[int]string и возвращает два среза: один с ключами, другой со значениями.
// Пример:
// Вход: map[int]string{1: "a", 2: "b"}
// Выход: []int{1, 2}, []string{"a", "b"}
func slice(mapa map[int]string) ([]int, []string) {
	nums := []int{}
	strings := []string{}
	for k, v := range mapa {
		nums = append(nums, k)
		strings = append(strings, v)
	}
	return nums, strings
}

func main() {
	m := map[int]string{1: "a", 2: "b"}
	fmt.Println(slice(m))
}
