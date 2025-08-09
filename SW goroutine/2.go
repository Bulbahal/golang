package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 2. Атомарный счетчик
// Создайте потокобезопасный счетчик, который может инкрементироваться и
// возвращать текущее значение, используя только атомарные операции (atomic).
type Counter struct {
	value int64
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}
func (c *Counter) GetValue() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	result := Counter{0}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			result.Increment()
		}()
	}
	wg.Wait()
	fmt.Println(result.GetValue())
}
