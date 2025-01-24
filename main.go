package main

import (
	"fmt"

	"github.com/gastonjarju/cli-calculator/input"
	"github.com/gastonjarju/cli-calculator/operations"
)

/*
Validate User Inputs:

Add checks to ensure the operator (userOperations) and numbers are valid. For instance, you can validate the operator as one of +, -, *, or /.
Error Handling:

Enhance error handling to provide feedback when invalid inputs are entered instead of simply returning.
*/
func main () {
	fmt.Println("Welcome to the Go CLI Calculator!")
	userInput := input.UserInputs()
	result := operations.PerformCalculation(userInput)
	fmt.Println("The result is:",result)

	
}