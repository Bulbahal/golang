package main

import (
	"fmt"
	"reflect"
)

// Задача 2: Изменение значений полей структуры
// Создайте структуру с несколькими полями.
// Напишите функцию, которая принимает указатель на эту структуру и изменяет значения полей, используя рефлексию.
// Например, если поле имеет тип `int`, увеличьте его на 10, если `string`, добавьте " (изменено)".
type Persona struct {
	Name string
	Age  int
}

func changeStruct(p interface{}) {
	val := reflect.ValueOf(p).Elem()
	if val.Kind() != reflect.Struct {
		panic("not a struct")
	}
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if !field.CanSet() {
			continue
		}
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field.SetInt(field.Int() + 10)
		case reflect.String:
			field.SetString(field.String() + " (изменено)")
		}
	}
}
func main() {
	p := &Persona{Name: "Tom", Age: 18}
	fmt.Println("До:", *p)
	changeStruct(p)
	fmt.Println("После:", *p)
}
