package main

import "fmt"

//Подсчет гласных в строке
//Дана строка. Посчитайте количество каждой гласной буквы (a, e, i, o, u).

func main() {
	text := "Hello, world!"
	m := map[string]int{}
	for _, v := range text {
		if string(v) == "a" || string(v) == "e" || string(v) == "i" || string(v) == "o" || string(v) == "u" {
			m[string(v)]++
		}
	}
	fmt.Println(m)
}
