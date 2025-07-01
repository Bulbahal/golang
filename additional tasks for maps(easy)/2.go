package main

import "fmt"

// Добавление элемента в мапу
// Добавьте новый ключ "orange" со значением 20 в мапу из предыдущей задачи и выведите обновлённую мапу.
func main() {
	m := map[string]int{
		"apple":  10,
		"banana": 10,
		"cherry": 15,
	}
	m["orange"] = 20
	fmt.Println(m)
}
