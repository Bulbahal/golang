package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CustomError наша кастомная ошибка для превышения лимита
type CustomError struct {
	message    string
	limit      int
	lastString string
}

// Error реализуем интерфейс error
func (e *CustomError) Error() string {
	return fmt.Sprintf("%s, limit: %d, last string: %s", e.message, e.limit, e.lastString)
}

// countStrings функция для подсчета строк с обработкой ошибок
func countStrings(filePath string, limit int) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0
	var lastLine string

	for scanner.Scan() {
		counter++
		lastLine = scanner.Text()

		if counter > limit {
			return counter, &CustomError{
				message:    "string count exceed limit",
				limit:      limit,
				lastString: lastLine,
			}
		}
	}

	// Обрабатываем возможные ошибки сканера
	if err := scanner.Err(); err != nil {
		if errors.Is(err, io.EOF) {
			return counter, nil
		}
		return counter, fmt.Errorf("scanner error: %w", err)
	}

	return counter, nil
}

func main() {
	baseDir := filepath.Join("files", "3")
	inputPath := filepath.Join(baseDir, "in1.txt")

	// Устанавливаем лимит строк
	limit := 150 // Произвольное значение

	// Получаем количество строк
	count, err := countStrings(inputPath, limit)
	if err != nil {
		var customErr *CustomError
		if errors.As(err, &customErr) {
			// Обрабатываем нашу кастомную ошибку
			fmt.Println("string count exceed limit, please read another file =) err:", err.Error())
		} else {
			// Обрабатываем другие ошибки
			fmt.Printf("Error counting strings: %v\n", err)
		}
		return
	}

	// Выводим результат если нет ошибок
	fmt.Printf("Total strings: %d\n", count)
}
