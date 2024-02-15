// StringMinusString.go
package Act

import "strings"

func StringMinusString(a string, b string) string {
	result := strings.ReplaceAll(a, b, "")
	return result
}
