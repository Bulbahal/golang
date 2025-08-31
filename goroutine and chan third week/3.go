package main

import (
	"fmt"
	"time"
)

//3. Буферизованный vs небуферизованный канал
//Сравните поведение:
//Небуферизованного канала (make(chan int)).
//Буферизованного (make(chan int, 3)).
//Отправьте 3 значения без чтения и посмотрите, где происходит блокировка.
//Цель: понять разницу между буферизованными и небуферизованными каналами.

func main() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 1
	}()
	time.Sleep(time.Millisecond * 100)
	fmt.Println(<-ch1)
	ch2 := make(chan int, 3)
	go func() {
		for i := 1; i <= 3; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	time.Sleep(time.Millisecond * 100)
	for v := range ch2 {
		fmt.Println(v)
	}
}
