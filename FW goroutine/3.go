package main

import (
	"fmt"
	"sync"
	"time"
)

// 3. Горутина с задержкой
// Создайте горутину, которая ждёт 2 секунды (time.Sleep), а затем выводит "Done". Главный поток должен дождаться её завершения.
func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Done")
	}()
	wg.Wait()
}
