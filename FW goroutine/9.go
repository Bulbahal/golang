package main

import (
	"fmt"
	"sync"
)

// 9. Паника в горутине
// Напишите программу, где одна из горутин вызывает panic, и обработайте её с помощью recover в главном потоке.
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Горутина упала, но мы ее перехватили: %v\n", r)
				}
			}()
			defer wg.Done()
			if i == 5 {
				panic("Одна из горутин равна пяти")
			}
		}()
	}
	wg.Wait()
	fmt.Println("Конец программы")
}
