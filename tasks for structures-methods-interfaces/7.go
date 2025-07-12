package main

import (
	"fmt"
	"sort"
)

// 7. Сортировка структур
// Задача: Создайте структуру Student (Name, Grade).
// Создайте слайс студентов и отсортируйте его по полю Grade (используя sort.Interface).
type Student struct {
	Name  string
	Grade int
}
type ByGrade []Student

func (a ByGrade) Len() int           { return len(a) }
func (a ByGrade) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByGrade) Less(i, j int) bool { return a[i].Grade < a[j].Grade }
func main() {
	people := ByGrade{
		{"Bob", 7},
		{"Alice", 10},
		{"Tom", 4},
	}
	sort.Sort(people)
	fmt.Println(people)
}
