package main

import (
	"fmt"
	"time"
)

//2. Множество горутин
//Запустите 5 горутин, каждая из которых выводит свой номер (от 1 до 5). Убедитесь, что все горутины успевают выполниться.

func main() {
	go func() { fmt.Println("I'm a go routine number 1") }()
	go func() { fmt.Println("I'm a go routine number 2") }()
	go func() { fmt.Println("I'm a go routine number 3") }()
	go func() { fmt.Println("I'm a go routine number 4") }()
	go func() { fmt.Println("I'm a go routine number 5") }()
	time.Sleep(2 * time.Second)
}
