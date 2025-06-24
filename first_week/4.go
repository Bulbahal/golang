package main

import "fmt"

// 4. Факториал числа
// Реализуйте функцию, которая вычисляет факториал заданного числа.
var a int = 5

func main() {
	result := 1
	for i := 1; i <= a; i++ {
		result *= i

	}
	fmt.Println(result)
}
