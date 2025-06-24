package main

import "fmt"

// 5. Обратная строка
// Напишите программу, которая принимает строку от пользователя и выводит ее в обратном порядке.
var text = "Hello World"

func main() {
	result := ""
	for i := len(text) - 1; i >= 0; i-- {
		result += string(text[i])
	}
	fmt.Println(result)
}
