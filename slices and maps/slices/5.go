package main

import "fmt"

// Передача среза по указателю
// Напишите функцию, которая принимает указатель на срез строк и добавляет строку "modified" в конец.
// Пример:
// Вход: &[]string{"a", "b"}
// Выход: ["a" "b" "modified"]
func add(slice *[]string) {
	*slice = append((*slice), "modified")
}

func main() {
	slice := []string{"a", "b", "c"}
	fmt.Println("Вход", &slice)
	add(&slice)
	fmt.Println("Выход", slice)
}
