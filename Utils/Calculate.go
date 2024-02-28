// Calculate.go
package Utils

import (
	"regexp"
	"strings"
)

func Calculate(input string) {
	operator := Operator(input)

	operands := strings.Split(input, operator)
	if len(operands) > 2 {
		panic("Неверный формат выражения, не более двух операндов в выражении\n\n\n\n\n")
		return
	}
	if len(operands) < 2 {
		panic("Неверный формат выражения, не менее двух операндов в выражении\n\n\n\n\n")
		return
	}

	var operand1, operand2 string

	if operands[0][0] == '"' && operands[0][len(operands[0])-1] == '"' {
		operands[0] = operands[0][1 : len(operands[0])-1]
		if len(operands[0]) <= 12 {
			operand1 = "str"
		} else {
			panic("Первый операнд слишком длинный, максимальная длина 10 символов\n\n\n\n\n")
		}
	} else {
		panic("Первый операнд должен быть в кавычках\n\n\n\n\n")
	}

	if operands[1][0] == '"' && operands[1][len(operands[1])-1] == '"' {
		operands[1] = operands[1][1 : len(operands[1])-1]
		if len(operands[1]) <= 12 {
			operand2 = "str"
		} else {
			panic("Второй операнд слишком длинный, максимальная длина 10 символов\n\n\n\n\n")
		}
	} else if matched, _ := regexp.MatchString("\\b([1-9]|10)\\b", operands[1]); matched {
		operand2 = "int"
	} else {
		panic("Второй операнд должен быть числом от 1 до 10 или строкой\n\n\n\n\n")
	}

	Switch(operand1, operand2, operator, operands)
}
