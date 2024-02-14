// main.go
package main

import (
	"fmt"
)

func main() {
	for {
		input := GetUserInput()
		fmt.Println("Введенные данные: ", input)
		if IsExitCommand(input) {
			HandleExit()
			return // Прерывание выполнения функции main(), что приведет к завершению программы
		}

		// Выполнение операции и вывод результата.
		result := Calculate(input)
		//	fmt.Println("Результат: ", result)
		fmt.Println("Результат: \"" + fmt.Sprint(result) + "\"")
	}
}
