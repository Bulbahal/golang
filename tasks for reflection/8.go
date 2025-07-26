package main

import (
	"fmt"
	"reflect"
)

// Задача 8: Подсчет значений полей
// Создайте структуру с несколькими полями, содержащими числовые значения.
// Напишите функцию, которая с помощью рефлексии подсчитывает сумму всех числовых полей и возвращает её.
type Calendar struct {
	Age    int
	Month  int
	Day    int
	Hour   int
	Minute int
	Second float64
}

func Sum(s interface{}) float64 {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	var sum float64
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			sum += float64(field.Int())
		case reflect.Float32, reflect.Float64:
			sum += field.Float()
		}
	}
	return sum
}
func main() {
	c := Calendar{2025, 7, 26, 22, 04, 56.23}
	fmt.Println(Sum(c))
}
