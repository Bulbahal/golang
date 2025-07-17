package main

// Задача 5: Игровая логика
// Создайте простую игру, где игрок может бросать кубик. Если выпадает число 1, вызывайте панику.
// Перехватите её и дайте игроку возможность продолжить игру.
// Подсказка: используйте метод из пакета math

import (
	"fmt"
	"math/rand"
)

func rollDice() int {
	return rand.Intn(6) + 1
}

type DiceError struct {
	Code int
}

func (d *DiceError) Error() string {
	return fmt.Sprintf("Вам выпала единица, означающая поражение: %d", d.Code)
}
func playDice() {
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					if d, ok := err.(*DiceError); ok {
						fmt.Printf("Паника перехвачена %v\n", d)
					} else {
						panic(err)
					}
				}
			}()
			dice := rollDice()
			fmt.Printf("Выпало: %v\n", dice)
			if dice == 1 {
				panic(&DiceError{Code: dice})
			}
		}()
		fmt.Println("Желаете продолжить?")
		fmt.Println("Введите да/нет")
		input := ""
		fmt.Scanln(&input)
		if input != "да" {
			break
		}
	}
}
func main() {
	fmt.Println("Начинаем игру")
	playDice()
	fmt.Println("Конец")
}
