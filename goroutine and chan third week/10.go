package main

import (
	"fmt"
	"time"
)

// 10. Pipeline из каналов
// Постройте конвейер (pipeline) из трёх этапов:
// Генератор чисел (1, 2, 3...).
// Умножение чисел на 2.
// Вывод результата.
// Каждый этап — отдельная goroutine, связанная каналами.
// Цель: освоить построение сложных цепочек обработки.
func generate(num ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range num {
			out <- n
			time.Sleep(100 * time.Millisecond)
		}
		close(out)
	}()
	return out
}
func multiple(in <-chan int, number int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			result := n * number
			out <- result
			time.Sleep(100 * time.Millisecond)
		}
		close(out)
	}()
	return out
}
func print(in <-chan int) {
	for n := range in {
		fmt.Printf("Result: %d\n", n)
	}
}
func main() {
	fmt.Println("start")
	numbers := generate(1, 2, 3, 4, 5)
	multiplied := multiple(numbers, 10)
	print(multiplied)
	fmt.Println("end")
}
