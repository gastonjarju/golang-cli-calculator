package operations

import (
	"errors"
	"fmt"

	"github.com/gastonjarju/cli-calculator/input"
)

func PerformCalculation (input input.CalculationInput) (float64, error) {
		var result float64
		
		switch input.Operation {
		case "+":
			result = input.Num1 + input.Num2
		case "-":
			result = input.Num1 - input.Num2
		case "/":
			if(input.Num2 == 0){
				return 0, errors.New("division by zero")
				} else {
					result = input.Num1 / input.Num2
			}
		case "*":
			result = input.Num1 * input.Num2
		default:
			fmt.Println("Invalid operation")
	
		}

	
	return result, nil
}