package main

import (
	"fmt"

	"github.com/gastonjarju/cli-calculator/input"
	"github.com/gastonjarju/cli-calculator/operations"
	"github.com/gastonjarju/cli-calculator/output"
)

func main () {
	fmt.Printf("Welcome to the Go CLI Calculator!\n")
	
	userInput := input.UserInputs()
		
	result, err := operations.PerformCalculation(userInput)
	if(err !=nil){
		fmt.Println("Error", err)
	} else {
		fmt.Println("The result is:",result)
	}

	result = output.LogResult(result)



	
}