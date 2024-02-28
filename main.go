// calc.go
package main

import "calculatorGo2/Utils"

func main() {
	for {
		input := Utils.GetUserInput()
		if Utils.IsExitCommand(input) {
			Utils.HandleExit()
			return
		}

		Utils.Calculate(input)
	}
}
