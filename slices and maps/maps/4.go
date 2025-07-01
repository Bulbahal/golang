package main

import (
	"fmt"
)

// Проверка на анаграммы
// Напишите функцию, которая проверяет, являются ли две строки анаграммами
// (содержат одинаковые символы в разном порядке), используя мапу для подсчета символов.
// Пример:
// Вход: "listen", "silent"
// Выход: true
func anagramma(s1 string, s2 string) bool {
	m1 := map[string]int{}
	m2 := map[string]int{}
	for _, v := range s1 {
		m1[string(v)]++
	}
	for _, v := range s2 {
		m2[string(v)]++
	}
	//return reflect.DeepEqual(m1, m2)
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(anagramma("listen", "silent"))
}
