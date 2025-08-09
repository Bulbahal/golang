package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

//4. Thread-safe булев флаг
//Реализуйте булев флаг, который можно безопасно устанавливать и проверять
//из разных горутин, используя только atomic.

func main() {
	var stopFlag atomic.Bool
	go func() {
		for !stopFlag.Load() {
			time.Sleep(1 * time.Second)
			fmt.Println(".")
		}
	}()
	fmt.Println(stopFlag)
	time.Sleep(6 * time.Second)
	stopFlag.Store(true)
	fmt.Println(stopFlag)
}
