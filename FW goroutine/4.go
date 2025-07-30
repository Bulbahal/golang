package main

import (
	"fmt"
	"sync"
	"time"
)

// 4. Горутины и цикл
// Запустите 3 горутины внутри цикла, передавая каждой значение счётчика (i). Убедитесь, что вывод соответствует ожиданиям (0, 1, 2).
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("n:", n)
			time.Sleep(1 * time.Second)
		}(i)
	}
	wg.Wait()
}
