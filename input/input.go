package input

import "fmt"

type CalculationInput struct {
	Operation string
	Num1 float64
	Num2 float64
}

func UserInputs() CalculationInput  {

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

func InputValidityCheck (inputCheck string) bool {
	if(inputCheck != "+" && inputCheck !="-" && inputCheck !="*" && inputCheck !="/") {
		fmt.Println("Invalid Operator.Try again")
		return false
	}

	return true
}