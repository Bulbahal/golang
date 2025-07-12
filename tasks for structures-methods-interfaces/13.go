package main

import (
	"fmt"
)

// 13. Реализация интерфейса error
// Задача:
// Создайте структуру DivError с полями dividend и divisor. Реализуйте для нее интерфейс error,
// чтобы при делении на 0 возвращалось сообщение: "деление на 0: dividend={X}, divisor=0".
type DivError struct {
	dividend, divisor float64
}

func (d *DivError) Error() string {
	return fmt.Sprintf("деление на 0: dividend={%f}, divisor=0", d.dividend)
}
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &DivError{dividend: a, divisor: 0}
	}
	return a / b, nil
}

func main() {
	a, err := Divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(a)
	}
}
