package main

import (
	"fmt"
	"time"
)

//1. Простой обмен сообщениями
//Создайте два goroutine:
//Один отправляет числа от 1 до 5 в канал.
//Второй читает их и выводит в консоль.
//Цель: понять базовую работу каналов.

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i <= 5; i++ {
			ch <- i
			fmt.Printf("Отправлено: %d\n", i)
		}
		close(ch)
	}()
	go func() {
		for nun := range ch {
			fmt.Printf("Получено: %d\n", nun)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	time.Sleep(4 * time.Second)
}
