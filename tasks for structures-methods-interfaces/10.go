package main

import (
	"errors"
	"fmt"
)

//10. Обработка ошибок в методах
//Задача: Создайте структуру Wallet с полем Balance. Добавьте метод Spend(amount float64) error,
//который уменьшает баланс, но возвращает ошибку, если денег недостаточно.

type Wallet struct {
	Balance float64
}

func (w *Wallet) Spend(amount float64) error {
	if w.Balance < amount {
		return errors.New("Wallet has insufficient funds")
	}
	w.Balance -= amount
	return nil
}
func (w *Wallet) ShowBalance() float64 {
	return w.Balance
}
func main() {
	balance := Wallet{156.2}
	err := balance.Spend(156.3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(balance.ShowBalance())
	}
}
