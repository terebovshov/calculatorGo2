package Utils

import (
	"calculatorGo2/Act"
	"fmt"
	"strconv"
	"strings"
)

func Switch(operand1, operand2, operator string, operands []string) {
	if operand1 == "str" && operand2 == "str" {
		Operand2String(operator, operands)
	} else if operand1 == "str" && operand2 == "int" {
		Operand2StringInt(operator, operands)
	}
}

func Operand2String(operator string, operands []string) {
	var result string
	switch operator {
	case " + ":
		result = Act.StringPlusString(operands[0], operands[1])
		fmt.Println("\"" + result + "\"")
	case " - ":
		if !strings.Contains(operands[0], operands[1]) {
			result = operands[0]
			fmt.Println("\"" + result + "\"")
		} else {
			result = Act.StringMinusString(operands[0], operands[1])
			fmt.Println("\"" + result + "\"")
		}
	default:
		panic("Между строкой и строкой работает только сложение или вычитание\n\n\n\n\n")
	}
}

func Operand2StringInt(operator string, operands []string) {
	n, err := strconv.Atoi(operands[1])
	if err != nil {
	}

	var result string

	switch operator {
	case " * ":
		result = Act.StringMultiplyInt(operands[0], n)
		if len(result) > 40 {
			result = result[:40]
			fmt.Println("\"" + result + "..." + "\"")
		} else {
			fmt.Println("\"" + result + "\"")
		}
	case " / ":
		result = Act.StringDivideInt(operands[0], n)
		if len(result) > 40 {
			result = result[:40]
			fmt.Println("\"" + result + "..." + "\"")
		} else {
			fmt.Println("\"" + result + "\"")
		}
	default:
		panic("Между строкой и числом работает только умножение или деление\n\n\n\n\n")
	}
}
