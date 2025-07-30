package main

import (
	"fmt"
	"sync"
)

//6. Исправление гонки через sync.Mutex
//Модифицируйте предыдущую задачу, используя sync.Mutex, чтобы избежать гонки данных.

func main() {
	i := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	for j := 0; j < 1000; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			i++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(i)
}
