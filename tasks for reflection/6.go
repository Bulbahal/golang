package main

import (
	"fmt"
	"reflect"
)

// Задача 6: Проверка тегов структуры
// Напишите программу, которая определяет структуру с полями, содержащими теги.
// Используйте рефлексию для получения значений тегов и выведите их на экран.
// Например, если у вас есть теги `json`, выведите, какие поля будут сериализоваться в JSON.
type Persoon struct {
	Name string `json:"name"`
	Age  int
}

func InspectTag(s interface{}) {
	v := reflect.TypeOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic("not a struct")
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if jsonTag, ok := field.Tag.Lookup("json"); ok {
			fmt.Printf("Поле %s будет сериализована как  %s\n", field.Name, jsonTag)
		}
	}
}
func main() {
	p := Persoon{"Иван", 18}
	InspectTag(p)
}
