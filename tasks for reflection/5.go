package main

import (
	"fmt"
	"reflect"
)

// Задача 5: Создание обобщенной функции
// Реализуйте обобщенную функцию, которая принимает срез любого типа и возвращает срез,
// содержащий только те элементы, которые соответствуют заданному условию (например, больше определённого значения).
// Используйте рефлексию для работы с различными типами.
func FilterSlices(slice interface{}, f func(interface{}) bool) interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic("FilterSlices() given a non-slice type")
	}
	result := reflect.MakeSlice(v.Type(), 0, 0)
	for i := 0; i < v.Len(); i++ {
		if f(v.Index(i).Interface()) {
			result = reflect.Append(result, v.Index(i))
		}
	}
	return result.Interface()
}
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	filtNum := FilterSlices(numbers, func(x interface{}) bool {
		return x.(int) >= 3
	})
	fmt.Println(filtNum)
	str := []string{"a", "b", "c", "d", "e"}
	filtStr := FilterSlices(str, func(x interface{}) bool {
		return x.(string) != "a"
	})
	fmt.Println(filtStr)
}
