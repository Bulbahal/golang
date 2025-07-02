package main

import (
	"fmt"
	"strings"
)

// 2. Подсчет частоты слов
// Дана строка с словами. Посчитайте, сколько раз каждое слово встречается в строке.
func main() {
	text := "Карл у Клары украл кораллы а Клара у Карла украла кларнет"
	result := strings.Split(text, " ")
	m := make(map[string]int)
	for _, v := range result {
		m[v]++
	}
	fmt.Println(m)
}
