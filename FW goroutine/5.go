package main

import (
	"fmt"
	"sync"
)

//5. Гонка данных
//Создайте программу с общей переменной, которую увеличивают несколько горутин. Покажите, что без синхронизации возможна гонка данных.

//	func main() {
//		i := 10
//		go func() {
//			i++
//			fmt.Println(i)
//
//		}()
//		go func() {
//			fmt.Println(i)
//			i++
//		}()
//		go func() {
//			fmt.Println(i)
//			i++
//		}()
//		time.Sleep(1 * time.Second)
//	}
func main() {
	i := 0
	var wg sync.WaitGroup
	for j := 0; j < 1000; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
		}()
	}
	wg.Wait()
	fmt.Println(i)
}
