package main

import (
	"fmt"
)

//3. Метод с получателем по значению и указателю
//Задача: Создайте структуру Counter с полем Value. Добавьте два метода:

type Counter struct {
	Value int
}

func (c Counter) Double() int {
	return c.Value * 2
}
func (c *Counter) Increase() {
	c.Value++
}
func main() {
	res := Counter{Value: 3}
	fmt.Println("Исходное значение:", res.Value)
	fmt.Println("После удвоения:", res.Double())
	fmt.Println("Значение Value после удвоения:", res.Value)
	res.Increase()
	fmt.Println("Значение Value после инкрементации(использование указателя):", res.Value)
}
