package main

//Напишите функцию, которая создает срез целых чисел от 1 до n,
//затем добавляет к нему числа от n+1 до 2n с помощью append.
//Выведите исходный и измененный срез.
//
//Пример:
//Вход: n = 3
//Выход:
//Исходный: [1 2 3]
//После append: [1 2 3 4 5 6]
import "fmt"

func slices(n int) []int {
	a := []int{}
	for i := 1; i <= n; i++ {
		a = append(a, i)
	}
	fmt.Println("Исходный:", a)
	//b := a
	for i := n + 1; i <= n*2; i++ {
		a = append(a, i)
	}
	return a
}
func main() {
	n := 3
	fmt.Println("После append: ", slices(n))
}
