package main

import "fmt"

//  9. Обобщенный контейнер:
//     Реализуйте обобщенный контейнер `Container[T any]`, который может хранить одно значение и предоставлять методы
//     для получения, установки и очистки этого значения.
type Container[T any] struct {
	value T
}

func (c *Container[T]) Get() {
	fmt.Println(c.value)
}
func (c *Container[T]) Set(NewValue T) {
	c.value = NewValue
}
func (c *Container[T]) Clear() {
	var zero T
	c.value = zero
}
func main() {
	var c Container[int]
	c = Container[int]{10}
	fmt.Println("Изначальное значение:", c.value)
	c.Set(100)
	fmt.Println("Значение после установки:", c.value)
	c.Clear()
	fmt.Println("Значение после отчистки:", c.value)
}
