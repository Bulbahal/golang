package main

import (
	"fmt"
	"math"
	"sync"
)

//8. Горутины и возврат значений
//Создайте функцию, которая запускает горутину, вычисляющую квадрат числа, и возвращает результат через разделяемую переменную (с синхронизацией).

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	number := 49
	var rs float64
	wg.Add(1)
	go func(n int) {
		defer wg.Done()
		sqrt := math.Sqrt(float64(n))
		mu.Lock()
		rs = sqrt
		mu.Unlock()
	}(number)
	wg.Wait()
	fmt.Printf("Квадратный корень из числа %d равен: %0.2f", number, rs)
}
