package main

import (
	"fmt"
	"time"
)

//16. Работа с временем в структуре
//Задача:
//Создайте структуру Event с полями Title и Time (типа time.Time).
//Добавьте метод IsAfter(now time.Time) bool, проверяющий, что событие еще не наступило.

type Event struct {
	Title string
	Time  time.Time
}

func (e *Event) IsAfter(now time.Time) bool {
	return e.Time.After(now)
}
func main() {
	eventTime := time.Now().Add(time.Hour)
	event := Event{Title: "Ревью", Time: eventTime}
	n := time.Now()
	if event.IsAfter(n) {
		fmt.Printf("%s ещё не началось\n", event.Title)
	} else {
		fmt.Printf("%s уже началось\n", event.Title)
	}
}
