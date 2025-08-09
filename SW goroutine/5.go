package main

import (
	"fmt"
	"sync"
	"time"
)

// 5. Банковский счет с RW Mutex
// Создайте структуру банковского счета, где:
// Проверка баланса происходит очень часто
// Изменение баланса (пополнение/списание) происходит редко
// Оптимизируйте с помощью sync.RWMutex.
type Bank struct {
	mu      sync.RWMutex
	balance int
}

func (b *Bank) GetBalance() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.balance
}
func (b *Bank) Deposit(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance += amount
}
func (b *Bank) Withdraw(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance -= amount
}
func main() {
	balance := &Bank{balance: 100}
	for i := 0; i < 10; i++ {
		go func(id int) {
			for {
				fmt.Println(id, ":", balance.GetBalance())
				time.Sleep(1 * time.Second)
			}
		}(i)

	}
	time.Sleep(2 * time.Second)
	fmt.Println("Пополнение баланса на 100")
	balance.Deposit(100)
	fmt.Println(balance.GetBalance())
	time.Sleep(2 * time.Second)
	fmt.Println("Забрали 50")
	balance.Withdraw(50)
	fmt.Println(balance.GetBalance())
	time.Sleep(2 * time.Second)
}
