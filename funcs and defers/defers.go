package main

import (
	"fmt"
	"time"
)

// 1. Базовый defer
// Напишите функцию, которая выводит "Start", затем откладывает вывод "Defer",
// а потом выводит "End". Объясните порядок вывода.
//func start(s string) string {
//	fmt.Println(s)
//	defer fmt.Println("Defer")
//	return "End"
//}
//func main() {
//	fmt.Println(start("Start"))
//}

// 2. Порядок выполнения defer (LIFO)
// Напишите функцию с несколькими defer и покажите, что они выполняются в обратном порядке (LIFO — Last In, First Out).
//func lifo(s string) string {
//	defer fmt.Println("defer 1")
//	defer fmt.Println("defer 2")
//	defer fmt.Println("defer 3")
//	return s
//}
//func main() {
//	fmt.Println(lifo("Start"))
//}

// 3. Defer и возвращаемые значения
// Напишите функцию, которая использует defer для изменения возвращаемого значения. Объясните результат.
//func deferWithReturn() (x int) {
//	defer func() { x++ }()
//	return 5
//}
//func main() {
//	fmt.Println(deferWithReturn())
//}

// 4. Defer и аргументы функции
// Напишите функцию, где defer использует переменную, которая изменяется после defer. Объясните результат.
//func deferuse() (x int) {
//	num := 5
//	defer func() { x++ }()
//	fmt.Println("num=", num)
//	return num
//}
//func main() {
//	fmt.Println(deferuse())
//}

// 5. Defer внутри цикла
// Напишите функцию с циклом, где defer вызывается на каждой итерации.
// Объясните, почему все defer выполняются после цикла.
//func deferFor() int {
//	num := 0
//	for i := 0; i < 10; i++ {
//		defer fmt.Println(num)
//	}
//	return num
//}
//func main() {
//	fmt.Println(deferFor())
//}

// 6. Defer и изменение именованного возвращаемого значения (в Go можно типу возвращаемого значения присвоить имя)
// Напишите функцию, где defer изменяет именованное возвращаемое значение после return.
//func names() (x string) {
//	name := "Magomed"
//	//defer fmt.Println(name)
//	defer func() { x = "Maga" }()
//	fmt.Println(name)
//	return name
//}
//func main() {
//	fmt.Println(names())
//}

//7. Defer и замыкание (closure)
//Напишите функцию, где defer использует замыкание для доступа к переменной после её изменения.
//Обьяснить что такое замыкание
//func main() {
//	deferClosure()
//}
//func deferClosure() {
//	x := 10
//	defer func() {
//		fmt.Println(x)
//	}()
//	x = 20
//	fmt.Println(x)
//	x = 30
//
//}

// 8. Defer и цепочка вызовов
// Напишите функцию с несколькими defer, где один defer вызывает другую функцию, которая тоже содержит defer.
// Покажите порядок выполнения.
//func main() {
//	fmt.Println("hello world")
//	oneDefer()
//}
//func oneDefer() {
//	fmt.Println("hello")
//	defer fmt.Println("oneDefer")
//	defer func() {
//		fmt.Println("world")
//		twoDefer()
//	}()
//	fmt.Println("end")
//}
//func twoDefer() {
//	fmt.Println("hello hello")
//	defer fmt.Println("world")
//	fmt.Println("goodbye")
//}

// 9. Defer и возврат значения в условном блоке
// Напишите функцию, где defer изменяет возвращаемое значение, но сам return находится внутри if.
// Покажите, как defer влияет на результат.
//func deferBlock() int {
//	num := 5
//	defer func() {
//		if num > 0 {
//			fmt.Println("num =", num+1)
//		}
//	}()
//	return num
//}
//func main() {
//	fmt.Println(deferBlock())
//}

// 10. Использование defer для измерения времени выполнения
// Напишите программу, которая измеряет время выполнения функции(в функции можно вызвать цикл на миллион итераций)
// Используйте defer для вывода времени, прошедшего с момента начала выполнения функции до её завершения.
func main() {
	measuringTime()
}
func measuringTime() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	for i := 0; i <= 1000000; i++ {

	}
}
