// main.go
package main

import (
	"fmt"
)

func main() {
	for {
		input := GetUserInput()
		if IsExitCommand(input) {
			HandleExit()
			return // Прерывание выполнения функции main(), что приведет к завершению программы
		}

		// Выполнение операции и вывод результата.
		result := Calculate(input)

		// Преобразование результата в строку
		resultStr := fmt.Sprint(result)

		// Проверка длины результата и его обработка.
		if len(resultStr) > 40 {
			// Обрезка ответа после 40 символов
			resultStr = resultStr[:40]
			// Добавление многоточия в конце
			resultStr += "..."
		}
		// Вывод результата.
		fmt.Println("Результат: \"" + resultStr + "\"")
	}
}
