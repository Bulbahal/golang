package main

import (
	"fmt"
	"sync"
	"time"
)

// 9. Хранилище временных меток
// Создайте потокобезопасное хранилище последних временных меток событий, где:
// Запись происходит редко (обновление метки)
// Чтение происходит часто
// Используйте sync.RWMutex.
type Date struct {
	mu    sync.RWMutex
	place string
	time  int
}

func NewDate(initPlace string, initTime int) *Date {
	return &Date{
		place: initPlace,
		time:  initTime,
	}
}
func (c *Date) Get() (string, int) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.place, c.time
}
func (c *Date) Set(value string, n int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.place = value
	c.time = n
}
func main() {
	date := NewDate("2010-01-01", time.Now().Second())
	go func() {
		for {
			time.Sleep(5 * time.Second)
			newVal, newNum := "time_after", time.Now().Second()
			date.Set(newVal, newNum)
			fmt.Println("Данные обновлены")

		}
	}()
	for i := 0; i < 100; i++ {
		go func(id int) {
			for {
				val, num := date.Get()
				fmt.Printf("Горутина %d: %s %d\n", id, val, num)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}
	time.Sleep(1 * time.Minute)
	fmt.Println("Конец программы")
}
