package main

import "fmt"

//14. Метод с variadic-аргументами
//Задача:
//Создайте структуру Calculator с методом Sum(numbers ...int) int, который возвращает сумму переданных чисел.

type Calculator struct{}

func (c *Calculator) Sum(numbers ...int) int {
	sun := 0
	for _, number := range numbers {
		sun += number
	}
	return sun
}

func main() {
	j := Calculator{}
	result := j.Sum(1, 2, 3)
	fmt.Println(result)
}
