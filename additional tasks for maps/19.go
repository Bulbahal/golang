package main

import (
	"fmt"
	"strings"
)

//19. Найти слово с максимальной частотой
//Дан текст. Найдите слово, которое встречается чаще всего (регистр не учитывать).

func periodicity(arr []string) string {
	m := make(map[string]int)
	index := 0
	result := ""
	for _, v := range arr {
		m[v]++
	}
	for k, v := range m {
		if v > index {
			index = v
			result = k
		}

	}
	return result
}

func main() {
	str := "hello world golang programming language hello hello"
	arr := strings.Split(str, " ")
	fmt.Println(periodicity(arr))
}
