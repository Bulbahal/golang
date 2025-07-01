package main

import "fmt"

// Сравнение длины и capacity
// Напишите программу, которая:
// Создает срез с помощью make([]int, 3, 5),
// Добавляет элементы, пока len не превысит cap,
// Выводит len и cap после каждого добавления.
// Пример вывода:
// Исходный: len=3, cap=5
// После добавления 4: len=4, cap=5
// После добавления 6: len=6, cap=10 (удвоение capacity)
func main() {
	zeus := make([]int, 3, 5)
	capa := cap(zeus)
	fmt.Println("Исходный: len =", len(zeus), "cap =", cap(zeus))
	for i := 4; len(zeus) <= cap(zeus); i++ {
		zeus = append(zeus, i)
		if capa < cap(zeus) {
			fmt.Println("После добавления ", i, ": len=", len(zeus), ", cap=", cap(zeus), " (удвоение capacity)")
			break
		}
		fmt.Println("После добавлении", i, ":", "len = ", len(zeus), "cap = ", cap(zeus))
	}
}
