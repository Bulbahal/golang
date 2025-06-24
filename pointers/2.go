package main

import "fmt"

//2. Увеличение числа через указатель
//Задача:
//Напишите функцию Increment(n *int), которая увеличивает значение переданного числа на 1.
//
//Пример:
//
//num := 5
//Increment(&num)
//fmt.Println(num) // 6
//Условие:
//
//Функция не должна возвращать значение, а изменять переданную переменную.

func Increment(n *int) {
	*n++
}
func main() {
	n := 5
	Increment(&n)
	fmt.Println(n)
}
