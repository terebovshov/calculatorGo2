// StringDivideInt.go
package Act

func StringDivideInt(a string, n int) string {
	if n <= 0 {
		panic("Делить можно только на целые числа от 1 до 10 включительно \n\n\n\n\n")
	}
	length := len(a)
	index := length / n
	result := a[:index]
	return result
}
