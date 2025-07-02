package main

//14. Найти индекс первого уникального символа
//Дана строка. Верните индекс первого символа, который встречается только один раз.

import "fmt"

func main() {
	str := "hhheello world"
	m := make(map[string]int)
	result := 0
	for _, v := range str {
		if string(v) != " " {
			m[string(v)]++
		}
	}
	for k, v := range str {
		if m[string(v)] == 1 {
			result = k
			break
		}

	}
	fmt.Println(result)
}
