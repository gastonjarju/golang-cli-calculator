package main

import (
	"fmt"

	"github.com/gastonjarju/cli-calculator/input"
	"github.com/gastonjarju/cli-calculator/operations"
)

func main () {
	fmt.Printf("Welcome to the Go CLI Calculator!\n")
	userInput := input.UserInputs()
		
	result := operations.PerformCalculation(userInput)
	fmt.Println("The result is:",result)

	
}