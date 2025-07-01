package main

import "fmt"

//Подсчет частоты элементов
//Напишите функцию, которая принимает срез строк и возвращает map[string]int, где ключ — это строка, а значение — количество её вхождений.
//Пример:
//Вход: []string{"a", "b", "a", "c", "b"}
//Выход: map[a:2 b:2 c:1]

func maps(slice []string) map[string]int {
	m := map[string]int{}
	for _, v := range slice {
		m[v]++
	}
	return m
}

func main() {
	slice := []string{"a", "b", "a", "c", "b"}
	fmt.Println(maps(slice))
}
