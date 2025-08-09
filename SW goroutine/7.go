package main

import (
	"fmt"
	"sync"
	"time"
)

// 7. Кэш статистики с редкими обновлениями
// Создайте кэш статистических данных, где:
// Данные обновляются раз в минуту (одна горутина)
// Данные читаются сотни раз в секунду (много горутин)
// Примените sync.RWMutex.
type CasheStat struct {
	mu  sync.RWMutex
	val string
	num int
}

func NewCasheStat(initVal string, initNum int) *CasheStat {
	return &CasheStat{
		val: initVal,
		num: initNum,
	}
}
func (c *CasheStat) Get() (string, int) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.val, c.num
}
func (c *CasheStat) Set(value string, n int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val = value
	c.num = n
}
func main() {
	cashe := NewCasheStat("Time", 1)
	go func() {
		for {
			time.Sleep(1 * time.Minute)
			newVal, newNum := "time_after", time.Now().Minute()
			cashe.Set(newVal, newNum)
			fmt.Println("Данные обновлены")

		}
	}()
	for i := 0; i < 100; i++ {
		go func(id int) {
			for {
				val, num := cashe.Get()
				fmt.Printf("Горутина %d: %s %d\n", id, val, num)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}
	time.Sleep(5 * time.Minute)
	fmt.Println("Конец программы")
}
