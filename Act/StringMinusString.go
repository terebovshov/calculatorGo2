// StringMinusString.go
package Act

import "strings"

func StringMinusString(a, b string) string {
	if !strings.Contains(a, b) {
		return a
	}

	result := strings.Replace(a, b, "", -1)
	return result
}
