// StringDivideInt.go
package Act

func StringDivideInt(s string, n int) string {
	//	if n <= 0 {panic("Деление на 0 или отрицательные числа запрещено")}
	length := len(s) / n
	result := s[:length]
	return result
}
