package main

import "fmt"

//9. Полиморфизм с интерфейсами
//Задача: Создайте интерфейс Animal с методом Sound() string. Реализуйте его для структур Dog и Cat.
//Затем создайте функцию MakeSound(animal Animal), которая вызывает Sound().

type Animal interface {
	Sound() string
}

type Dog struct {
	Name string
}

func (d Dog) Sound() string {
	return fmt.Sprintf("%s: Woof", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Sound() string {
	return fmt.Sprintf("%s: Meow", c.Name)
}

func MakeSound(animal Animal) {
	fmt.Println(animal.Sound())
}
func main() {
	d := Dog{"Spike"}
	c := Cat{"Tom"}
	MakeSound(d)
	MakeSound(c)
}
