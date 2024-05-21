package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Calculator function to perform operations based on the operand
func Calculator(value1, value2 float64, operand string) (float64, error) {
	// TODO: Implement Calculator function
	// use switch case for addition,substraction,multiplication and division
	switch operand {
	case "+":
		return value1 + value2, nil
	case "-":
		return value1 - value2, nil
	case "*":
		return value1 * value2, nil
	case "/":
		if value2 == 0 {
			return 0, errors.New("can't divide by zero")
		}
		return value1 / value2, nil
	default:
		return 0, errors.New("invalid")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the expression (e.g., 10 + 20): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	// TODO: Check if exact 3 parts are given. If not, throw an error
	//checked by using the length of parts not equal to 3 because index value from 0 to 2
	if len(parts) != 3 {
		fmt.Println("invalid input")
	}
	// TODO: Take all 3 values and pass to calculator function
	//take value1 and err then convert into float64
	value1, err := strconv.ParseFloat(parts[0], 64)
	//handle the error if value1 not a number
	if err != nil {
		fmt.Printf("Invalid value1: %v", err)
		return
	}
	//take value2 and err then convert into float64
	value2, err := strconv.ParseFloat(parts[2], 64)
	//handle the error if value2 not a number
	if err != nil {
		fmt.Printf("Invalid value2: %v", err)
		return
	}
	//take part[1] index which is operand
	operand := parts[1]
	//call the calculator function store  it in result variable
	result, err := Calculator(value1, value2, operand)
	//handle the error by calculated value get error
	if err != nil {
		fmt.Printf("error calculate in result: %v", err)
		return
	}
	// TODO: Print results
	fmt.Printf("Result: %v\n", result)
}
