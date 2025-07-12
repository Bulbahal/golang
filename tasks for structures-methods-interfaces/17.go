package main

import (
	"fmt"
	"sort"
)

//17. Интерфейс для сортировки по разным полям
//Задача:
//Создайте структуру Product (Name, Price, Rating). Реализуйте интерфейс sort.Interface
//для сортировки слайса []Product по любому из полей (выбор поля — через вложенную структуру-компаратор).

type Product struct {
	Name   string
	Price  int
	Rating int
}
type ByName []Product

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type ByPrice []Product

func (b ByPrice) Len() int           { return len(b) }
func (b ByPrice) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByPrice) Less(i, j int) bool { return b[i].Price < b[j].Price }

type ByRating []Product

func (c ByRating) Len() int           { return len(c) }
func (c ByRating) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByRating) Less(i, j int) bool { return c[i].Rating < c[j].Rating }

func main() {
	products := []Product{
		{"Chips", 3, 60},
		{"Coca-Cola", 1, 100},
		{"Ice-creame", 2, 90},
	}
	sort.Sort(ByName(products))
	fmt.Println(products)
	sort.Sort(ByPrice(products))
	fmt.Println(products)
	sort.Sort(ByRating(products))
	fmt.Println(products)
}
