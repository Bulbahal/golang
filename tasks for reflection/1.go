package main

import (
	"fmt"
	"reflect"
)

// Задача 1: Получение информации о типе
// Напишите программу, которая принимает переменную любого типа и использует рефлексию для получения и вывода следующей информации:
// - Имя типа
// - Размер типа в байтах
// - Наличие метода с определённым именем
type Stringer interface {
	String() string
}

func interfacing(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Println("Имя типа:", t.Name())
	fmt.Println("Размер типа в байтах:", t.Size())
	_, hasMethod := t.MethodByName("String")
	fmt.Println("Наличие метода String():", hasMethod)
	if _, ok := v.(string); ok {
		fmt.Println("Реализуйте метод Stringer")
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d лет)", p.Name, p.Age)
}
func main() {
	interfacing(54)
	interfacing(struct{}{})
	p := Person{"Tom", 42}
	interfacing(p)
	if s, ok := interface{}(p).(Stringer); ok {
		fmt.Println(s.String())
	}
}
