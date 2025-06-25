package main

// Создайте новую директорию с файлом main.go. Напишите код, в котором:
// нужно объединить слайс с выходными днями и слайс с рабочими в один слайс. Выведите на экран итоговый слайс с днями.
// Создайте новый коммит и отправьте его в удалённый репозиторий
import "fmt"

func main() {
	weekends := []string{"Saturday", "Sunday"}
	workdays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	workdays = append(workdays, weekends...)
	fmt.Println(workdays)
}
