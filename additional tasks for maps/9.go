package main

import (
	"fmt"
)

//Сгруппировать слова по длине
//Дан список слов. Сгруппируйте их в хешмапу.
//где ключ — длина слова, а значение — список слов такой длины.

func main() {
	//str := "Ехал Грека через реку видит Грека  в реке рак Сунул Грека руку в реку рак за руку Греку цап"
	//str1 := strings.Split(str, " ")
	str1 := []string{"hello", "world", "meow", "golang", "programming"}
	m := map[int][]string{}
	for _, v := range str1 {
		m[len(v)] = append(m[len(v)], v)
	}
	fmt.Println(m)
}
