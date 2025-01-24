package operations

import (
	"fmt"

	"github.com/gastonjarju/cli-calculator/input"
)

func PerformCalculation (input input.CalculationInput) float64 {
		var result float64
		switch input.Operation {
		case "+":
			result = input.Num1 + input.Num2
		case "-":
			result = input.Num1 - input.Num2
		case "/":
			// TODO: Take care of edge cases with the divide using an if-statement
			result = input.Num1 - input.Num2
		case "*":
			result = input.Num1 - input.Num2
		default:
			fmt.Println("Invalid operation")
	
		}
	return result
}