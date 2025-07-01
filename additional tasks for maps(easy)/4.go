package main

import "fmt"

// 4. Удаление элемента из мапы
// Удалите ключ "banana" из мапы fruits := map[string]int{"apple": 5, "banana": 10, "cherry": 15} и выведите результат.
func main() {
	fruits := map[string]int{"apple": 5, "banana": 10, "cherry": 15}
	delete(fruits, "banana")
	fmt.Println(fruits)
}
