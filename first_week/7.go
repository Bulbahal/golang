package main

import (
	"fmt"
	"strings"
)

// 7. Проверка палиндрома
// Напишите функцию, которая проверяет, является ли строка палиндромом.
var a = "Level"
var text = strings.ToLower(a)

func reverse() string {

	result := ""
	for i := len(text) - 1; i >= 0; i-- {
		result += string(text[i])
	}
	return result
}
func main() {
	if reverse() == text {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not palindrome")
	}
}
