package main

import "fmt"

//Фильтрация мапы по значению
//Напишите функцию, которая принимает map[string]int и возвращает новую мапу, содержащую только элементы со значениями больше заданного числа n.
//Пример:
//Вход: map[string]int{"a": 10, "b": 5, "c": 20}, n = 15
//Выход: map[c:20]

func examination(mapa map[string]int, num int) map[string]int {
	m := make(map[string]int)
	for k, v := range mapa {
		if v >= num {
			m[k] = v
		}
	}
	return m
}

func main() {
	m1 := map[string]int{"a": 10, "b": 5, "c": 20, "d": 30, "e": 40}
	n := 15
	fmt.Println(examination(m1, n))
}
