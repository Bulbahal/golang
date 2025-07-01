package main

import "fmt"

// Вложенные мапы (словарь словарей)
// Создайте map[string]map[string]int, где внешний ключ — имя студента, а внутренняя мапа хранит предметы и оценки.
// Напишите функцию для добавления оценки студенту.
// Пример:
// Вход:
//
// db := map[string]map[string]int{}
// addGrade(db, "Alice", "Math", 90)
// Выход: map[Alice:map[Math:90]]
func addGrade(db map[string]map[string]int, student, subject string, grade int) {
	if _, exist := db[student]; exist {
		db[student] = make(map[string]int)
		db[student][subject] = grade
		fmt.Println(db)
	} else {
		db[student] = make(map[string]int)
		db[student][subject] = grade
	}
}
func main() {
	db := map[string]map[string]int{
		"Alice": map[string]int{
			"Math": 0,
		},
		"Evgeniy": map[string]int{
			"Chemistry": 70,
		},
	}
	addGrade(db, "Егор", "Physics", 80)
	addGrade(db, "Alice", "Math", 90)
}
