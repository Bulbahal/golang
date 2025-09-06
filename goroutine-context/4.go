package main

// 4. Построение конвейеров обработки данных
// Реализуйте конвейер (pipeline), где:
// Генератор создаёт поток данных (например, числа от 1 до 100)
// Фильтр пропускает только данные, удовлетворяющие условию (например, чётные числа)
// Обработчик преобразует данные (например, умножает на 10)
// Все этапы должны поддерживать отмену через context.Context
// При ошибке на любом этапе весь конвейер останавливается

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func generator(ctx context.Context, n int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			select {
			case <-ctx.Done():
				return
			case out <- i:
				time.Sleep(20 * time.Millisecond)
			}
		}
	}()

	return out
}

func filter(ctx context.Context, in <-chan int) (<-chan int, <-chan error) {
	out := make(chan int)
	errc := make(chan error, 1)

	go func() {
		defer close(out)
		for v := range in {
			if v == 42 {
				errc <- errors.New("нашли плохое число 42")
				return
			}

			if v%2 == 0 {
				select {
				case <-ctx.Done():
					return
				case out <- v:
				}
			}
		}
		errc <- nil
	}()

	return out, errc
}

func processor(ctx context.Context, in <-chan int) (<-chan int, <-chan error) {
	out := make(chan int)
	errc := make(chan error, 1)

	go func() {
		defer close(out)
		for v := range in {
			select {
			case <-ctx.Done():
				return
			case out <- v * 10:
			}
		}
		errc <- nil
	}()

	return out, errc
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nums := generator(ctx, 100)

	filtered, ferr := filter(ctx, nums)

	processed, perr := processor(ctx, filtered)

	for v := range processed {
		fmt.Println("Result:", v)
	}

	if err := <-ferr; err != nil {
		fmt.Println("Ошибка фильтра:", err)
		cancel()
	}
	if err := <-perr; err != nil {
		fmt.Println("Ошибка обработчика:", err)
		cancel()
	}

	fmt.Println("Конвейер завершён ✅")
}
