package internal

import "fmt"

func PrintMainMenu() {
	fmt.Print(`Word Trainer

1. Показать словарь
2. Добавить слово
3. Удалить слово
4. Начать тренировку
5. Повторить ошибки
6. Показать статистику
7. Справка
0. Выход
`)
}
