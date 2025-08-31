package main

import (
	"fmt"
	"time"
)

//8. Канал только для чтения/записи
//Создайте функцию, которая:
//Принимает канал только для чтения (<-chan int).
//Возвращает канал только для записи (chan<- int).
//Цель: понять, как ограничивать направление каналов.

func reader(ch <-chan int) {
	for v := range ch {
		fmt.Println("reader ", v)
	}
}
func writer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
func main() {
	ch := make(chan int, 5)
	go writer(ch)
	go reader(ch)
	time.Sleep(1 * time.Second)
}
