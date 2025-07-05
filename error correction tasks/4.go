package main

import (
	"errors"
	"fmt"
)

// 4. Цепочка вызовов с ошибками: Реализуйте цепочку из 3 функций,
// где каждая может вернуть ошибку, и обрабатывайте ошибки на верхнем уровне.
func resulting(f func() (int, error)) {
	result, err := f()
	if err != nil {
		fmt.Println("Ошибка : ", err.Error())
	} else {
		fmt.Println(result)
	}
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Деление на ноль")
	}
	return a / b, nil
}

func Multiply(a, b int) (int, error) {
	if b == 0 || a == 0 {
		return 0, errors.New("На ноль умножать нет смысла")
	}
	return a * b, nil
}

func Sum(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("Одно из значений отрицательное")
	}
	return a + b, nil
}

func main() {
	resulting(func() (int, error) {
		return Divide(3, 0)
	})
	resulting(func() (int, error) {
		return Multiply(3, 0)
	})
	resulting(func() (int, error) {
		return Sum(-3, 0)
	})
}
