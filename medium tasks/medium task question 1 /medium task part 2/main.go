package main

import (
	"fmt"
	"math"
)

// Создайте новую директорию и новый файл main.go. Напишите код, в котором:
// объявите два новых типа AmericanVelocity и EuropeanVelocity;
// выполните преобразование скорости 120.4 м/сек в км/ч и присвойте результат переменной с типом EuropeanVelocity;
// выполните преобразование скорости 130 м/с в миль/ч и присвойте результат переменной с типом AmericanVelocity;
// примечание: 1 миля = 1.609 км. Если потребуется, округлите значение до 2 знаков после запятой, для округления обратитесь к пакету math.
// Создайте новый коммит с вашими решениями задач и отправьте в удалённый репозиторий
type AmericanVelocity float64
type EuropeanVelocity float64

func main() {
	var sec AmericanVelocity = 120.4
	var mil EuropeanVelocity = 130
	sec = sec * 3.6
	mil = mil * 1.609
	fmt.Printf("%.2f\n", sec)
	fmt.Printf("%.2f\n", mil)
	hr := sec * 100
	hrrs := math.Round(float64(hr))
	km := mil * 100
	kmrs := math.Round(float64(km))
	fmt.Println(hrrs / 100)
	fmt.Println(kmrs / 100)
}
