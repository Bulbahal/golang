package main

import (
	"fmt"
	"time"
)

// 6. Таймауты в каналах
// Используйте select с time.After, чтобы:
// Ждать данных из канала не более 2 секунд.
// Если данных нет — вывести "Timeout!".
// Цель: реализовать таймауты для операций с каналами.
func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 500)
		}
		close(ch)
	}()
	timeout := time.After(time.Second * 2)
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println("Получено сообщений: ", msg)
		case <-timeout:
			fmt.Println("Timeout!")
			return
		}
	}
}
