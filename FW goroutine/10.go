package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// 10. Завершение горутин по флагу
// Создайте горутину, которая работает в бесконечном цикле, пока не будет установлен флаг (через атомарную переменную atomic.Bool)
func main() {
	var stopFlag atomic.Bool
	go func() {
		for !stopFlag.Load() {
			fmt.Println("work goroutine")
			time.Sleep(1 * time.Second)
		}
		fmt.Println("goroutine stopped")
	}()
	time.Sleep(2 * time.Second)
	stopFlag.Store(true)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("end programm")
}
