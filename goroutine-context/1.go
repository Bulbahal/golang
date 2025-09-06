package main

import (
	"context"
	"fmt"
	"time"
)

//1. Таймауты и отмена операций
//Напишите функцию, которая:
//Выполняет долгую операцию (например, time.Sleep или цикл вычислений)
//Принимает context.Context для контроля выполнения
//Прекращает работу, если контекст отменён (например, по таймауту 2 секунды)
//Возвращает ошибку, если операция не завершилась вовремя

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	done := make(chan error, 1)
	go func() {
		time.Sleep(5 * time.Second)
		done <- nil
	}()
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case err := <-done:
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("done")
		}
	}
}
