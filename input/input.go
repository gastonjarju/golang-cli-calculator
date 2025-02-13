package input

import (
	"errors"
	"fmt"
)

type CalculationInput struct {
	Operation string
	Num1 float64
	Num2 float64
}

func UserInputs() CalculationInput  {

	var userOperations string 
	var firstNum, secondNum float64

	
	for {
			fmt.Println("Enter operator: +, -, /, *:")
			fmt.Scanln(&userOperations)
		
			_, err := OperationsValidityCheck(userOperations)

			if(err !=nil){
				fmt.Println("Error:", err)
			continue

		}
		break
	} 

	for {
		fmt.Println("Enter first number:")
		_, err := fmt.Scanln(&firstNum) 
		if(err != nil) {
			fmt.Println("Invalid input. Please enter a valid number")
			fmt.Scanln()
			continue
		}
		
		_, err = NumberValidityCheck(firstNum)
		
		if(err != nil) {
			fmt.Println("Error:", err)
			continue
		}
		break
	}
	
	
	for {
		fmt.Println("Enter second number:")
		_, err := fmt.Scanln(&secondNum)
		if(err != nil) {
				fmt.Println("Invalid input. Please enter a valid number")
				fmt.Scanln()
				continue

			}

			_, err = NumberValidityCheck(secondNum)
				if(err != nil) {
					fmt.Println("Error:", err)
					continue
				}
				break

		}

	return CalculationInput{
		Operation: userOperations,
		Num1: firstNum,
		Num2: secondNum,
	}
}

func OperationsValidityCheck (inputCheck string) (bool, error) {
	if(inputCheck != "+" && inputCheck !="-" && inputCheck !="*" && inputCheck !="/") {
		return false, errors.New("invalid Operator.Try again")
	}

	return true, nil
}

func NumberValidityCheck (inputNum float64) (bool, error) {
	if(inputNum < 0 || inputNum > 1e9){
		return false, errors.New("invalid number inserted. Try again")
	}

	return true, nil
}