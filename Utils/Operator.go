package Utils

import (
	"strings"
)

func Operator(input string) string {
	var operator string
	operators := [...]string{" + ", " - ", " * ", " / "}
	for _, op := range operators {
		if strings.Contains(input, op) {
			operator = op
			break
		} else if op == operators[len(operators)-1] {
			panic("Оператор не найден, строка не является математической операцией\n\n\n\n\n")
		}
	}
	return operator
}
