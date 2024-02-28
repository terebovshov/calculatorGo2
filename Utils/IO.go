// IO.go = Input/Output.go
package Utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput() string {
	fmt.Print("\nВозможные операции в калькуляторе:\n[\"str\" - \"str\"] [\"str\" - \"str\"] [\"str\" * int] [\"str\" / int]\n(или 'x' для выхода)\n\nВведите выражение:\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func IsExitCommand(input string) bool {
	return strings.ToLower(input) == "x"
}

func HandleExit() {
	fmt.Println("\nВыход из программы.Cпасибо за использование! \n\n")
}
