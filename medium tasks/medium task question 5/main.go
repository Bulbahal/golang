package main

import "fmt"

//Напишите функцию contains, которая принимает на вход два параметра: слайс строк a и строку x.
//Функция должна проверять, содержится ли строка x в слайсе a, и возвращать булево значение.
//Создайте вариативную функцию getMax, которая находит максимальное целое число из переданных на вход параметров.
//Выведите на экран результат вызова функций.
//Создайте коммит с вашим решением и отправьте ветку удалённый репозиторий

func constains(slice []string, x string) bool {
	for _, str := range slice {
		if x == str {
			return true
		}
	}
	return false

}
func getMax(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
func main() {
	fmt.Println(getMax(1, 2, 30, 4, 5))
	var meow = []string{"Кирилл", "Андрей", "Павел", "Мастурбек"}
	fmt.Println(constains(meow, "Мастурбек"))
}
