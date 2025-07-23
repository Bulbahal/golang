package main

import "fmt"

//  5. Обобщенные карты:
//     Реализуйте обобщенную карту `Map[K comparable, V any]` с методами `Set(key K, value V)` и `Get(key K) V`.
type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Set(key K, value V) {
	m[key] = value
}
func (m Map[K, V]) Get(key K) V {
	return m[key]
}
func main() {
	m1 := make(Map[string, int])
	m1.Set("a", 1)
	m1.Set("b", 2)
	fmt.Println(m1.Get("a"))
	fmt.Println(m1.Get("b"))
	m2 := make(Map[string, string])
	m2.Set("hello", "world")
	m2.Set("coca", "cola")
	fmt.Println(m2.Get("hello"))
	fmt.Println(m2.Get("coca"))
}
