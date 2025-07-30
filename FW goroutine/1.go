package main

import (
	"fmt"
	"time"
)

// 1. Простая горутина
// Напишите программу, которая запускает горутину, выводящую "Hello, Go!",
// и дожидается её завершения с помощью time.Sleep.
func main() {
	go fmt.Println("Hello World")
	go func() {
		fmt.Println("Hello World")
	}()
	time.Sleep(1 * time.Second)
}
