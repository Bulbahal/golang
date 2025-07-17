package main

import "fmt"

//Задача 1: Обработка паники
//Создайте программу, которая вызывает функцию doSomething, в которой происходит паника.
//Используйте recover для перехвата паники и вывода сообщения об ошибке.

func main() {
	n := 1
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("hello world")
	doSomething(n)
	fmt.Println("done")
}
func doSomething(n int) {
	if n > 1 {
		panic("doSomething: n > 1")
	}
	fmt.Println("doSomething:", n)
}
