package main

import (
	"fmt"
	"time"
)

// 5. select и мультиплексирование
// Напишите программу с двумя каналами:
// Один отправляет числа каждые 500 мс.
// Второй — каждые 1 сек.
// Используйте select, чтобы читать из того канала, который готов.
// Цель: освоить мультиплексирование каналов.
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 1; ; i++ {
			time.Sleep(time.Millisecond * 500)
			ch1 <- i
		}
	}()
	go func() {
		for i := 10; ; i++ {
			time.Sleep(time.Second * 1)
			ch2 <- i
		}
	}()
	for i := 0; i < 10; i++ {
		select {
		case num1 := <-ch1:
			fmt.Printf("Быстрый канал: %v\n", num1)
		case num2 := <-ch2:
			fmt.Printf("Медленный канал: %v\n", num2)
		}
	}
}
