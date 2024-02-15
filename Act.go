// Act.go
// 		StringPlusString
// 		StringMinusString
//		StringMultiplyInt
//		StringDivideInt

package main

import "strings"

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
