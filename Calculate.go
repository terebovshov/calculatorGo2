// Calculate.go
package main

import (
	"strconv"
	"strings"
)

// Calculate выполняет математическую операцию и возвращает результат.
func Calculate(calc string) interface{} {
	args := strings.Fields(calc) // Разбиение строки на аргументы.

	if len(args) != 3 { // Проверка корректности количества аргументов.
		panic("некорректный формат ввода, ожидалось: 'аргумент1 оператор аргумент2'")
	}

	// Извлечение аргументов
	operand1 := args[0]
	operator := args[1]
	operand2 := args[2]

	// Проверим, находится ли operand1 в двойных кавычках
	// Если да - удалим их для обработки значения, если нет - паника.
	if !strings.HasPrefix(operand1, `"`) || !strings.HasSuffix(operand1, `"`) {
		panic("operand1 должно быть в кавычках")
	} else {
		operand1 = strings.Trim(args[0], `"`) // Удаляем двойные кавычки из строки args[0].
	}

	//	Функции Check:
	//	Check.CheckOutputLength(calc)

	// Выполнение операции в зависимости от оператора.
	switch operator {
	case "+":
		if !strings.HasPrefix(operand2, `"`) || !strings.HasSuffix(operand2, `"`) {
			panic("operand2 должно быть в кавычках") // Вызываем панику, если строка не находится в кавычках.
		} else {
			operand2 = strings.Trim(args[2], `"`) // Удаляем двойные кавычки из строки args[2].
		}
		return StringPlusString(operand1, operand2) // Сложение строк.
	case "-":
		if !strings.HasPrefix(operand2, `"`) || !strings.HasSuffix(operand2, `"`) {
			panic("operand2 должно быть в кавычках") // Вызываем панику, если строка не находится в кавычках.
		} else {
			operand2 = strings.Trim(args[2], `"`) // Удаляем двойные кавычки из строки args[2].
		}
		return StringMinusString(operand1, operand2) // Вычитание строк.
	case "/":
		n, _ := strconv.Atoi(operand2) // конвертируем второй операнд в число
		if n < 1 || n > 10 {           // если ошибка или число не в диапазоне от 1 до 10
			panic("недопустимое число для деления") // вызываем panic
		}
		return StringDivideInt(operand1, n) // вызываем функцию для деления строки на число
	case "*":
		n, _ := strconv.Atoi(operand2) // конвертируем второй операнд в число
		if n < 1 || n > 10 {           // если ошибка или число не в диапазоне от 1 до 10
			panic("недопустимое число для умножения") // вызываем panic
		}
		return StringMultiplyInt(operand1, n) // вызываем функцию для деления строки на число
	default:
		// Паника при неверном операторе.
		panic("неверный оператор: " + operator)
	}
}

func StringPlusString(a string, b string) string {
	return a + b
}

func StringMinusString(a string, b string) string {
	result := strings.ReplaceAll(a, b, "")
	return result
}

func StringMultiplyInt(operand1 string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += operand1
	}
	return result
}

func StringDivideInt(s string, n int) string {
	//	if n <= 0 {panic("Деление на 0 или отрицательные числа запрещено")}
	length := len(s) / n
	result := s[:length]
	return result
}
