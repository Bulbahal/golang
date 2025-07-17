package main

import (
	"fmt"
)

// Задача 2: Пользовательская паника
// Напишите программу, которая создает пользовательский тип ошибки и вызывает панику с этим типом.
// В основной функции перехватите эту панику и выведите сообщение с деталями ошибки.
type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}
func rs(meow int) {
	if meow == 0 {
		panic(&MyError{
			Code: 500,
			Msg:  "Число равняется нулю",
		})
	}

}
func main() {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(*MyError); ok {
				fmt.Printf("Паника перехвачена %v\n", err)
			} else {
				fmt.Printf("Неизвестная паника:", r)
			}
		}
	}()
	rs(0)
	fmt.Println("Ошибок не было")
}
