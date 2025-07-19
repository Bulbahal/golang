package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//скопируйте in.txt в директорию data рядом с файлом main.go.
//Напишите программу, которая считывает данные из файла, проверяет, что все поля заполнены и записывает считанные данные в файл data/data_out.txt в формате Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n.
//Если какое-то поле не заполнено, то программа должна вызвать панику, передав строку вида parse error: empty field on string %d.
//Объявите необходимые отложенные вызовы.
//Обработайте панику таким образом, чтобы после обработки на экран вывелось содержимое файла data_out.txt, которое было записано до возникновения паники.

func main() {
	// Определяем пути к файлам
	baseDir := filepath.Join("files", "1")
	inputPath := filepath.Join(baseDir, "in.txt")
	outputDir := filepath.Join(baseDir, "data")
	outputPath := filepath.Join(outputDir, "data_out.txt")

	// Создаем директорию data если ее нет
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		panic(fmt.Sprintf("Ошибка создания директории: %v", err))
	}

	// Открываем входной файл
	inputFile, err := os.Open(inputPath)
	if err != nil {
		panic(fmt.Sprintf("Ошибка открытия файла %s: %v", inputPath, err))
	}
	defer inputFile.Close()

	// Создаем выходной файл
	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(fmt.Sprintf("Ошибка создания файла %s: %v", outputPath, err))
	}
	defer outputFile.Close()

	// Обработчик паники
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Программа завершилась с ошибкой:", r)

			// Читаем что успели записать
			content, err := os.ReadFile(outputPath)
			if err != nil {
				fmt.Println("Не удалось прочитать выходной файл:", err)
				return
			}

			fmt.Println("Содержимое data_out.txt до ошибки:")
			fmt.Println(string(content))
		}
	}()

	// Обработка файла
	scanner := bufio.NewScanner(inputFile)
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		// Разбиваем строку по вертикальной черте |
		fields := strings.Split(line, "|")

		// Проверяем что ровно 3 поля (в вашем формате)
		if len(fields) != 3 {
			panic(fmt.Sprintf("Ошибка: должно быть 3 поля разделенных |, строка %d", lineNumber))
		}

		// Проверяем что нет пустых полей
		for i, field := range fields {
			fields[i] = strings.TrimSpace(field)
			if fields[i] == "" {
				panic(fmt.Sprintf("Ошибка: пустое поле в строке %d", lineNumber))
			}
		}

		// Форматируем строку для вывода (в вашем формате)
		output := fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n",
			lineNumber, fields[0], fields[1], fields[2])

		// Записываем в выходной файл
		if _, err := outputFile.WriteString(output); err != nil {
			panic(fmt.Sprintf("Ошибка записи: %v", err))
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Ошибка чтения файла: %v", err))
	}

	fmt.Println("Файл успешно обработан! Результат в:", outputPath)
}
