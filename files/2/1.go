package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

//скопируйте in1.txt в директорию data рядом с файлом main.go.
//Напишите программу, которая построчно считывает файл.
//Выведите в консоль количество строк в формате Total strings: %d.
//Корректно обработайте в отложенном вызове ошибки закрытия файловых дескрипторов.
//Корректно обработайте ошибку окончания файла EOF

func main() {
	// 1. Определяем пути к файлам
	baseDir := filepath.Join("files", "2")         // Изменили на files/2
	inputPath := filepath.Join(baseDir, "in1.txt") // Теперь in1.txt

	// 2. Открываем входной файл
	inputFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v\n", err)
		return
	}
	defer func() {
		// 3. Обработка ошибки закрытия файла в отложенном вызове
		if err := inputFile.Close(); err != nil {
			fmt.Printf("Ошибка закрытия файла: %v\n", err)
		}
	}()

	// 4. Создаем счетчик строк
	lineCount := 0

	// 5. Читаем файл построчно
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lineCount++
	}

	// 6. Проверяем ошибки сканера
	if err := scanner.Err(); err != nil {
		// 7. Специальная проверка на EOF (хотя для Scanner она не нужна)
		if err != io.EOF {
			fmt.Printf("Ошибка чтения файла: %v\n", err)
			return
		}
	}

	// 8. Выводим результат
	fmt.Printf("Total strings: %d\n", lineCount)
}
