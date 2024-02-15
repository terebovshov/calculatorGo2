// StringMultiplyInt.go
package main

func StringMultiplyInt(operand1 string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += operand1
	}
	return result
}
