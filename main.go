package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"word-trainer/internal"
)

func main() {
	// Обрабатываем аргументы
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "Ошибка: обнаружены лишние аргументы.")
		os.Exit(1)
	}
	if len(args) != 0 && args[0] != "--help" && !strings.HasPrefix(args[0], "--data") {
		fmt.Fprintln(os.Stderr, "Ошибка: неизвестный аргумент.")
		os.Exit(1)
	}

	// Записываем путь к файлу
	var pathToJson string
	if len(args) != 0 && strings.HasPrefix(args[0], "--data") {
		pathToJson = args[0][7:]
	} else {
		pathToJson = "trainer-data.json"
	}

	// Проверяем на корректность пути к json-файлу
	if filepath.Ext(pathToJson) != ".json" {
		fmt.Fprintln(os.Stderr, "Ошибка: файл должен иметь расширение .json.")
		os.Exit(1)
	}
	info, err := os.Stat(pathToJson)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Fprintln(os.Stderr, "Ошибка: файл по указанному пути не найден.")
		os.Exit(1)
	} else if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при проверке файла:", err)
		os.Exit(1)
	}

	if !info.Mode().IsRegular() {
		fmt.Fprintln(os.Stderr, "Ошибка: указанный путь не является обычным файлом.")
		os.Exit(1)
	}

	// Загружаем содержимое json-файла
	jsonContent, err := os.ReadFile(pathToJson)
	if err != nil {
		fmt.Println("Ошибка чтения файла: ", err)
		return
	}

	// Выводим главное меню и обрабатываем ее
	scanner := bufio.NewScanner(os.Stdin) // Обработка идет так, ибо при повторных вводах через fmt.Scan() неправильные вводы остаются в буфере
	var choice int
	for {
		fmt.Print("Выберите действие: ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Ошибка чтения: ", err)
			}
			os.Exit(1)
		}

		choice, err = strconv.Atoi(strings.TrimSpace(scanner.Text())) // Превращаем строку в число

		if err != nil {
			fmt.Println("Ошибка: введите целое число от 0 до 7.")
			continue
		}

		if choice < 0 || choice > 7 {
			fmt.Println("Ошибка: введите число от 0 до 7.")
			continue
		}
		break
	}

	switch choice {
	case 1:
		internal.LoadData(jsonContent)
	case 0:
		return
	}
}
