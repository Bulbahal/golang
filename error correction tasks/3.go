package main

import (
	"fmt"
	"strconv"
)

//3. Проверка типа ошибки: Напишите код, который пытается преобразовать строку
//в число и по-разному обрабатывает ошибки парсинга и ошибки диапазона.

type ParsingError struct{}

func (p *ParsingError) Error() string {
	return "ParsingError"
}

type DiapazonError struct{}

func (d *DiapazonError) Error() string {
	return "DiapazonError"
}
func Conv(str string) (int, error) {
	result, error := strconv.Atoi(str)
	if error != nil {
		return 0, &ParsingError{}
	}
	if result < 0 {
		return 0, &DiapazonError{}
	}
	return result, nil
}
func main() {
	fmt.Println(Conv("qwef"))
	fmt.Println(Conv("-124"))
	fmt.Println(Conv("flps"))
	fmt.Println(Conv("420"))
}
