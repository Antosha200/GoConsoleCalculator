package main

import (
	"calculator/internal/calculator"
	"fmt"
	"log"
)

func main() {
	var a, b float64
	var operator string

	fmt.Print("Enter first number: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		log.Fatalf("Invalid input for first number: %v", err)
	}

	fmt.Print("Enter second number: ")
	_, err = fmt.Scan(&b)
	if err != nil {
		log.Fatalf("Invalid input for second number: %v", err)
	}

	for {
		fmt.Print("Enter operator (+, -, *, /, %, ^): ")
		_, err = fmt.Scan(&operator)
		if err != nil {
			log.Fatalf("Invalid input for operator: %v", err)
		}
		if isValidOperator(operator) {
			break
		} else {
			fmt.Println("Invalid operator. Please enter a valid operator (+, -, *, /, %, ^).")
		}
	}

	bc := calculator.BasicCalculator{}

	result, err := bc.Compute(a, b, operator)
	if err != nil {
		log.Fatalf("Error while computing: %v", err)
	}

	if result == float64(int(result)) {
		fmt.Printf("Result: %d\n", int(result))
	} else {
		fmt.Printf("Result: %.4f\n", result)
	}
}

func isValidOperator(operator string) bool {
	validOperators := []string{"+", "-", "*", "/", "%", "^"}
	for _, validOp := range validOperators {
		if operator == validOp {
			return true
		}
	}
	return false
}
