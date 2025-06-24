package main

//Создайте новую директорию с файлом main.go. Напишите код, в котором:
//объявите две переменные, первая — строка со значением 104, вторая — целое число со значением 35;
//приведите строку к целому числу, а целое число — к строке;

import (
	"fmt"
	"strconv"
)

func main() {
	str := "104"
	number, _ := strconv.Atoi(str)
	internet := 35
	fiveG := strconv.Itoa(internet)
	fmt.Printf("%T\n%T\n", fiveG, number)

}
