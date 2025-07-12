package main

import "fmt"

//15. Интерфейс для клонирования объектов
//Задача:
//Создайте интерфейс Cloner с методом Clone() Cloner.
//Реализуйте его для структуры Document (поле Content string), чтобы метод возвращал копию объекта.

type Cloner interface {
	Clone() Cloner
}
type Document struct {
	Content string
}

func (d *Document) Clone() Cloner {
	return &Document{Content: d.Content}
}
func main() {
	orig := &Document{Content: "Original Content"}
	cl := orig.Clone().(*Document)
	cl.Content = "Cloned"
	fmt.Println("Original Content:", orig.Content)
	fmt.Println("Cloned Content:", cl.Content)
}
