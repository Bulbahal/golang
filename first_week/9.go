package main

import "fmt"

//9. Числа Фибоначчи
//Напишите функцию, которая генерирует первые N чисел Фибоначчи.

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
func main() {
	fmt.Println(fibonacci(9))
}
