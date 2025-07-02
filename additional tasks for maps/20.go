package main

import (
	"fmt"
	"unicode"
)

// 20. Проверить, является ли строка панграммой
// Панграмма — это строка, содержащая все буквы алфавита.
// Проверьте, является ли данная строка панграммой (можно игнорировать регистр).
func IsPangram(str string) bool {
	m := make(map[rune]bool)
	for _, c := range str {
		if unicode.IsLetter(c) {
			m[unicode.ToLower(c)] = true
		}
	}
	return len(m) == 26
}

func main() {
	str := "pack my box with five dozen liquor jugs"
	fmt.Println(IsPangram(str))
	fmt.Println(IsPangram("fewijfioewjfweifjwejfiweffefefef"))
}
