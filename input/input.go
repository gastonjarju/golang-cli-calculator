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
			fmt.Scan(&userOperations)
		
			_, err := OperationsValidityCheck(userOperations)

			if(err !=nil){
				fmt.Println("Error:", err)
			continue

		}
		break
	} 

	for {
		fmt.Println("Enter first number:")
		fmt.Scan(&firstNum)

		_, err := NumberValidityCheck(int(firstNum))
			if(err != nil) {
				fmt.Println("Error:", err)
				continue
			}
			break
		}
		
		
		for {
			fmt.Println("Enter second number:")
			fmt.Scan(&secondNum)
			_, err := NumberValidityCheck(int(secondNum))
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

func NumberValidityCheck (inputNum int) (bool, error) {
	if(inputNum != 2){
		return false, errors.New("invalid number inserted. Try again")
	}

	return true, nil
}