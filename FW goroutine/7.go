package main

import (
	"fmt"
	"sync"
)

//7. Ожидание горутин с sync.WaitGroup
//Запустите 4 горутины, каждая из которых увеличивает счётчик, и дождитесь их завершения с помощью sync.WaitGroup.

func main() {
	i := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(4)
	for j := 0; j < 4; j++ {
		go func(g int) {
			defer wg.Done()
			mu.Lock()
			i++
			mu.Unlock()
			fmt.Printf("goroutine %d is done\n", g)
		}(j)
	}
	wg.Wait()
	fmt.Println("Итоговое число:", i)
}
