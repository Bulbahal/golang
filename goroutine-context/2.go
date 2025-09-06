package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

// 2. Параллельная обработка со сбором ошибок
// Реализуйте функцию, которая:
// Принимает слайс данных (чисел, строк и т.п.)
// Обрабатывает каждый элемент параллельно в горутинах
// Использует errgroup для управления горутинами
// Возвращает:
// Результаты обработки (если все элементы обработаны успешно)
// Первую возникшую ошибку (если хотя бы одна операция завершилась с ошибкой)
func main() {
	nums := []int{1, 2, 4, 5}
	result := make([]int, len(nums))
	var g errgroup.Group
	for i, n := range nums {
		i, n := i, n
		g.Go(func() error {
			if n == 3 {
				return errors.New("err on num 3")
			}
			result[i] = n * 2
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Result:", result)
}
