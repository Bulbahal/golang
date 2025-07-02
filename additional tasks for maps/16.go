package main

import "fmt"

// 16. Проверить, можно ли получить строку перестановкой другой
// Даны две строки. Проверьте, можно ли получить одну из другой перестановкой символов.
func isEqueal(s1, s2 string) bool {
	m1 := make(map[string]int)
	m2 := make(map[string]int)
	for _, v := range s1 {
		m1[string(v)]++
	}
	for _, v := range s2 {
		m2[string(v)]++
	}
	return len(m1) == len(m2)
}

func main() {
	str1 := "listen"
	str2 := "silent"
	fmt.Println(isEqueal(str1, str2))
}
