// StringDivideInt.go
package Act

func StringDivideInt(a string, n int) string {
	if n <= 0 {
		panic("Деление на 0 или отрицательные числа запрещено\n\n\n\n\n")
	}
	length := len(a)
	index := length / n
	result := a[:index]
	return result
}
