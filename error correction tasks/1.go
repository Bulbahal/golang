package main

import (
	"errors"
	"fmt"
)

//1. Напишите функцию, которая делит два числа и возвращает ошибку, если делитель равен нулю.

func delenie(a, b int) (int, error) {
	if b == 0 {
		err := errors.New("Деление на ноль")
		return 0, err
	}
	return a / b, nil
}
func main() {
	_, err := delenie(3, 0)
	fmt.Println(delenie(3, 3))
	if err != nil {
		fmt.Println("Ошибка ", err)
	}
}
