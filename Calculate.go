// Calculate.go
package main

import (
	"strconv"
	"strings"
)

// Calculate выполняет математическую операцию и возвращает результат.
func Calculate(calc string) interface{} {
	// Разбиение строки на аргументы и дальнейшая их обработка
	args := strings.Fields(calc)

	// Проверка длины каждого операнда
	for _, arg := range args {
		if len(arg) > 12 {
			panic("операнд слишком длинный, максимальная длина 10 символов")
		}
	}

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
