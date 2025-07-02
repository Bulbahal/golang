package main

import (
	"fmt"
	"strings"
)

// 15. Подсчет количества подстрок с одинаковыми символами на концах
// Дана строка. Посчитайте количество подстрок, которые начинаются и заканчиваются одним и тем же символом.
func substring(s []string) map[string]int {
	m := make(map[string]int)
	for _, v := range s {
		b := fmt.Sprintf("%s-%s", v[:1], v[len(v)-1:])
		m[b]++
		fmt.Println(b)
	}
	return m

}
func main() {
	str := "hello world"
	str1 := strings.Split(str, " ")
	fmt.Println(substring(str1))
}
