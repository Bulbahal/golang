package main

import (
	"fmt"
	"time"
)

// 2. Синхронизация горутин
// Напишите программу, где:
// Главная goroutine ждёт завершения двух рабочих goroutine с помощью канала.
// Рабочие goroutine отправляют сигнал (done <- true) после выполнения задачи.
// Цель: научиться синхронизировать горутины.
func main() {
	done := make(chan bool, 2)
	go func() {
		fmt.Println("Горутина номер один начала работу")
		time.Sleep(time.Second * 1)
		fmt.Println("Горутина номер один завершила работу")
		done <- true
	}()
	go func() {
		fmt.Println("Горутина номер два начала работу")
		time.Sleep(time.Second * 2)
		fmt.Println("Горутина номер два завершила работу")
		done <- true
	}()
	fmt.Println("Главная горутина ждёт завершение всех работ")
	<-done
	<-done
	fmt.Println("Все прошло удачно, всё сработало, конец программы")
}
