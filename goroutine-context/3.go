package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// 3. Управление жизненным циклом горутин
// Создайте программу, которая:
// Запускает N рабочих горутин (например, 3)
// Каждая горутина периодически выполняет свою задачу
// При получении сигнала SIGINT (Ctrl+C):
// Корректно останавливает все горутины
// Даёт им время на завершение текущих операций
// Выводит статус завершения
func main() {
	const workerCount = 3
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i < workerCount; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	cancel()
	wg.Wait()
	fmt.Println("All goroutines terminated")
}
func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker", id, "terminated")
			return
		default:
			fmt.Println("worker", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
