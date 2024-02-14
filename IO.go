// IO.go = Input/Output.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetUserInput получает ввод от пользователя.
func GetUserInput() string {
	fmt.Print("Введите выражение: [\"qwe\" + \"rty\"] [\"asd\" - \"s\"] [\"xyz\" * 3] [\"were\" / 2] (или нажмите 'q' для выхода): ") // Вывод приглашения для ввода пользователю
	scanner := bufio.NewScanner(os.Stdin)                                                                                             // Инициализация сканера для считывания ввода
	scanner.Scan()                                                                                                                    // Считывание введенной строки
	return scanner.Text()                                                                                                             // Возвращение введенной строки
}

// IsExitCommand проверяет, является ли строка командой выхода.
func IsExitCommand(input string) bool {
	return strings.ToLower(input) == "q" // Проверка наличия команды выхода в строке в нижнем регистре
}

// HandleExit обрабатывает выход из программы.
func HandleExit() {
	fmt.Println("Выход из программы. Cпасибо за использование!") // Вывод сообщения о завершении программы
}
