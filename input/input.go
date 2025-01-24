package input

import "fmt"

type CalculationInput struct {
	Operation string
	Num1 float64
	Num2 float64
}

// TODO: Add checks for non-numerics catch
// TODO: Add checks for non operation checks

func UserInputs() CalculationInput {

	var userOperations string 
	var firstNum, secondNum float64

	
	fmt.Println("Enter operator: +, -, /, *:")
	fmt.Scan(&userOperations)
	
	fmt.Println("Enter first number:")
	fmt.Scan(&firstNum)
	
	fmt.Println("Enter second number:")
	fmt.Scan(&secondNum)


	return CalculationInput{
		Operation: userOperations,
		Num1: firstNum,
		Num2: secondNum,
	}
}