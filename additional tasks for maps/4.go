package main

import "fmt"

//Найти первый неповторяющийся символ
//Дана строка. Найдите первый символ, который встречается только один раз.

func main() {
	str := "hhheello world"
	m := make(map[string]int)
	result := ""
	for _, v := range str {
		if string(v) != " " {
			m[string(v)]++
		}
	}
	for _, v := range str {
		if m[string(v)] == 1 {
			result = string(v)
			break
		}

	}
	fmt.Println(result)
}
