package main

import (
	"fmt"
	"reflect"
)

// Задача 7: Обработка интерфейсов
// Создайте несколько типов, реализующих один и тот же интерфейс.
// Напишите функцию, которая принимает интерфейс и с помощью рефлексии выводит
// информацию о конкретном типе, который реализует этот интерфейс.
type Speaker interface {
	Sound() string
}
type Dog struct{ Name string }

func (d Dog) Sound() string {
	return "Гав"
}

type Cat struct{ Name string }

func (c Cat) Sound() string {
	return "Мяу"
}
func inspectInterface(i interface{}) {
	v := reflect.ValueOf(i)
	t := v.Type()
	fmt.Println("Конкретный тип:", t.Name())
	fmt.Println("Кол-во методов:", t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		fmt.Println("Method:", t.Method(i).Name)
	}
}
func main() {
	d := Dog{Name: "Шарик"}
	c := Cat{Name: "Кот в сапогах"}
	inspectInterface(d)
	inspectInterface(c)
}
