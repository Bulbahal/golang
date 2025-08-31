package main

import (
	"fmt"
	"time"
)

// 9. Остановка goroutine через канал
// Реализуйте управление goroutine:
// Создайте канал stopChan.
// Goroutine работает в бесконечном цикле, но завершается при получении сигнала из stopChan.
// Цель: научиться останавливать горутины безопасно.
func worker(stopChan <-chan bool, id int) {
	for {
		select {
		case <-stopChan:
			fmt.Println("worker stopped", id)
			return
		default:
			fmt.Println("worker running", id)
			time.Sleep(time.Second)
		}
	}
}
func main() {
	stopChan := make(chan bool)
	for i := 0; i < 10; i++ {
		go worker(stopChan, i)
	}
	time.Sleep(2 * time.Second)
	close(stopChan)
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}
