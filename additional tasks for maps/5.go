package main

import (
	"fmt"
)

// Проверка анаграмм
// Даны две строки. Проверьте, являются ли они анаграммами (состоят из одних и тех же символов в разном порядке).
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
