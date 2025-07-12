package main

import (
	"fmt"
	"math"
)

// 11. Интерфейс для геометрических фигур с периметром
// Задача:
// Расширьте интерфейс Shape из задачи 2, добавив метод Perimeter() float64. Реализуйте его для Rectangle и Circle.
type Shape3 interface {
	Area3() float64
	Perimeter() float64
}
type Rectangle3 struct {
	Width, Height float64
}
type Circle3 struct {
	Radius float64
}

func (r *Rectangle3) Area3() float64 {
	return r.Width * r.Height
}
func (c *Circle3) Area3() float64 {
	return c.Radius * c.Radius * math.Pi
}
func (r *Rectangle3) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}
func (c *Circle3) Perimeter() float64 {
	return (math.Pi * c.Radius) * 2
}
func printingaut(s Shape3) {
	fmt.Printf("Фигура: %T\n", s)
	fmt.Printf("Площадь: %.2f\n", s.Area3())
	fmt.Printf("Периметр: %.2f\n", s.Perimeter())
	fmt.Println("-----")
}
func main() {
	a := &Circle3{Radius: 5}
	b := &Rectangle3{5, 5}
	printingaut(a)
	printingaut(b)
}
