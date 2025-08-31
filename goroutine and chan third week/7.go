package main

import (
	"fmt"
	"time"
)

// 7. Паттерн Worker Pool
// Создайте пул из 3 рабочих goroutine.
// Главная goroutine отправляет задачи (числа) в канал, рабочие обрабатывают их (например, умножают на 2).
// Цель: разобрать популярный шаблон Worker Pool.
func workerr(in <-chan int) {
	for v := range in {
		fmt.Println(v * 2)
	}
}
func main() {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go workerr(ch)
	}
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	time.Sleep(1 * time.Second)
}
