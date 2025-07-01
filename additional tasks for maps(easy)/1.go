package main

import "fmt"

// Создание и вывод мапы
// Создайте мапу map[string]int с ключами "apple", "banana", "cherry" и значениями 5, 10, 15. Выведите её.
func main() {
	m := map[string]int{
		"apple":  5,
		"banana": 10,
		"cherry": 15,
	}
	fmt.Println(m)
}
