package main

import "fmt"

//18. Динамическая диспетчеризация методов
//Задача:
//Создайте интерфейс Transport с методом Move(). Реализуйте его для структур Car, Bicycle и Plane.
//Напишите функцию StartRace(transports []Transport), которая вызывает Move() для всех участников.

type Transport interface {
	Move()
}
type Car struct {
	Name string
}

func (c *Car) Move() {
	fmt.Printf("%s мчит \n", c.Name)
}

type Bicecle struct {
	Name string
}

func (b *Bicecle) Move() {
	fmt.Printf("%s едет \n", b.Name)
}

type Plane struct {
	Name string
}

func (p *Plane) Move() {
	fmt.Printf("%s летит \n", p.Name)
}

func StartRace(transport []Transport) {
	for _, t := range transport {
		t.Move()
	}
}
func main() {
	transport := []Transport{
		&Car{"Audi"},
		&Bicecle{"Polygon"},
		&Plane{"Boeing"},
	}
	StartRace(transport)
}
