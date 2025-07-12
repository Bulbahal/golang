package main

import "fmt"

// 1. Простая структура и метод
// Задача: Создайте структуру Rectangle с полями Width и Height. Добавьте метод Area(),
// который вычисляет площадь прямоугольника.
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func main() {
	rec := Rectangle{width: 10, height: 5}
	fmt.Println(rec.Area())
}
