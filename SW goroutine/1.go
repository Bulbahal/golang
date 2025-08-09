package main

import (
	"fmt"
	"sync"
	"time"
)

//1. Кэш с RW Mutex
//Реализуйте потокобезопасный кэш, где операции чтения происходят значительно чаще, чем записи.
//Используйте sync.RWMutex для оптимизации производительности.

type Cashe struct {
	mu    sync.RWMutex
	items map[string]string
}

func NewCashe() *Cashe {
	return &Cashe{items: make(map[string]string)}
}
func (c *Cashe) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.items[key]
	return val, ok
}
func (c *Cashe) Set(key string, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value
}
func (c *Cashe) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

func main() {
	cashe := NewCashe()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		cashe.Set("Name", "Nikolai")
	}()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			value, ok := cashe.Get("Name")
			fmt.Printf("Горутина %d: name=%v, ok=%v\n", id, value, ok)
		}(i)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		cashe.Delete("Name")
	}()
	wg.Wait()
	fmt.Println("Конец программы")
}

//Вывод проверки на гонку данных:
//bulbahal@macbookair SW goroutine % go run -race 1.go
//Горутина 1: name=Nikolai, ok=true
//Горутина 3: name=Nikolai, ok=true
//Горутина 2: name=Nikolai, ok=true
//Горутина 4: name=Nikolai, ok=true
//Горутина 0: name=Nikolai, ok=true
//Конец программы
