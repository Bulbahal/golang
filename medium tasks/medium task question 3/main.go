package main

import "fmt"

// Создайте новую директорию. В ней создайте файл main.go. Напишите код, в котором:
// из слайса с днями недели надо скопировать в новый слайс рабочие дни, а из исходного слайса удалить скопированные, чтобы остались только выходные дни.
// выведите на экран слайсы с выходными и рабочими днями недели.
func main() {
	input := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	workDays := make([]string, len(input))
	copy(workDays, input[:5])
	fmt.Println(workDays)
	input = input[5:]
	fmt.Println(input)
}
