package main

import (
	"fmt"
	"reflect"
	"sort"
)

//Проверить, являются ли строки изоморфными
//Две строки называются изоморфными, если можно заменить символы одной строки на символы другой (с сохранением порядка).
//Например, "egg" и "add" изоморфны.

func isomorphic(s string, t string) bool {

	m1 := map[string]int{}
	m2 := map[string]int{}
	result1 := []int{}
	result2 := []int{}
	for _, v := range s {
		m1[string(v)]++
	}
	for _, v := range t {
		m2[string(v)]++
	}
	for _, v := range m1 {
		result1 = append(result1, v)
	}
	for _, v := range m2 {
		result2 = append(result2, v)
	}
	sort.SliceStable(result1, func(i, j int) bool {
		return result1[i] < result1[j]
	})
	return reflect.DeepEqual(result1, result2)
}
func main() {
	fmt.Println(isomorphic("egg", "add"))
	fmt.Println(isomorphic("dog", "cat"))
	fmt.Println(isomorphic("egg", "meow"))
	fmt.Println(isomorphic("kide", "tide"))
	fmt.Println(isomorphic("мука", "луна"))
	fmt.Println(isomorphic("банан", "тапок"))

}
