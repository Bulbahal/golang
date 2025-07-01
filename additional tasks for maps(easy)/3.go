package main

import "fmt"

// Проверка наличия ключа
// Дана мапа ages := map[string]int{"Alice": 25, "Bob": 30}.
// Проверьте, есть ли в ней ключ "Bob", и выведите "Exists" или "Not exists".
func main() {
	ages := map[string]int{"Alice": 25, "Bob": 30}
	if ages["Bob"] != 0 {
		fmt.Println("Exists")
	} else {
		fmt.Println("Not exists")
	}
	_, ok := ages["Alice"]
	if ok {
		fmt.Println("Exists")
	} else {
		fmt.Println("Not exists")
	}
}
