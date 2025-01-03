package main

import (
	"fmt"
	"os"
)

// Function to perform addition
func add(a, b float64) float64 {
	return a + b
}

// Function to perform subtraction
func subtract(a, b float64) float64 {
	return a - b
}

// Function to perform multiplication
func multiply(a, b float64) float64 {
	return a * b
}

// Function to perform division
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	var operation string
	var num1, num2 float64

	fmt.Println("Simple Calculator")
	fmt.Println("----------------")
	fmt.Println("Operations available: +, -, *, /")
	
	for {
		// Get the operation from user
		fmt.Print("\nEnter operation (+, -, *, /) or 'q' to quit: ")
		fmt.Scan(&operation)

		// Check if user wants to quit
		if operation == "q" {
			fmt.Println("Thanks for using the calculator!")
			os.Exit(0)
		}

		// Get the numbers from user
		fmt.Print("Enter first number: ")
		fmt.Scan(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scan(&num2)

		// Perform the calculation based on operation
		switch operation {
		case "+":
			result := add(num1, num2)
			fmt.Printf("Result: %.2f\n", result)
		
		case "-":
			result := subtract(num1, num2)
			fmt.Printf("Result: %.2f\n", result)
		
		case "*":
			result := multiply(num1, num2)
			fmt.Printf("Result: %.2f\n", result)
		
		case "/":
			result, err := divide(num1, num2)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("Result: %.2f\n", result)
			}
		
		default:
			fmt.Println("Invalid operation! Please use +, -, *, or /")
		}
	}
}