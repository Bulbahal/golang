package main

import (
	"fmt"
	"math"
)

// 2. Создание кастомных ошибок: Создайте тип NegativeNumberError и напишите функцию,
// которая вычисляет квадратный корень и возвращает эту ошибку, если число отрицательное.
type NegativeNumberError struct{}

func (e NegativeNumberError) Error() string {
	return "ошибка, число отрицательное"
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, NegativeNumberError{}
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(float64(-48)))
}
