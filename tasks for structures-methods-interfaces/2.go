package main

import (
	"fmt"
	"math"
)

// 2. Интерфейс "Фигура"
// Задача: Создайте интерфейс Shape с методом Area() float64.
// Реализуйте этот интерфейс для структур Circle (с полем Radius) и Rectangle (из первой задачи).
type Rectangle2 struct {
	width, height float64
}
type Circle struct {
	Radius float64
}
type Shape interface {
	Area2() float64
}

func (r Rectangle2) Area2() float64 {
	return r.width * r.height
}
func (c Circle) Area2() float64 {
	return math.Pi * c.Radius * c.Radius
}
func printArea(s Shape) {
	fmt.Println("Площадь:", math.Round(s.Area2()))
}
func main() {
	rect := Rectangle2{width: 10, height: 5}
	circle := Circle{Radius: 4}
	printArea(rect)
	printArea(circle)
}
