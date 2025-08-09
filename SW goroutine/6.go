package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//6. Атомарный генератор ID
//Реализуйте генератор уникальных идентификаторов, который гарантированно возвращает
//уникальные числа даже при вызове из множества горутин. Используйте atomic.

func main() {
	var start atomic.Int32
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rs := start.Add(1)
			fmt.Println(rs)
		}()
	}
	wg.Wait()
}
