package main

import "fmt"

// 2. Обобщенный стек:
// Реализуйте обобщенный стек `Stack[T any]` с методами `Push(value T)` и `Pop() T`,
// которые добавляют и удаляют элементы из стека соответственно.
type Stack[T any] struct {
	elements []T
}

func (p *Stack[T]) Push(element T) {
	p.elements = append(p.elements, element)
}
func (p *Stack[T]) Pop() T {
	p.elements = p.elements[:len(p.elements)-1]
	return p.elements[len(p.elements)-1]
}
func main() {
	myStack := Stack[string]{}
	myStack.Push("hello")
	myStack.Push("world")
	fmt.Println(myStack.Pop())
	myIntStack := Stack[int]{}
	myIntStack.Push(1)
	myIntStack.Push(2)
	fmt.Println(myIntStack.Pop())
}
