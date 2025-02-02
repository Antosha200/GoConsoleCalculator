package main

import (
	"calculator/internal/calculator"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a, b float64
	var operator string

	for {
		fmt.Print("Enter first number (or type 'exit' to quit): ")
		var input string
		_, err := fmt.Scan(&input)

		if err != nil {
			log.Fatalf("Invalid input: %v", err)
		}

		if input == "exit" {
			fmt.Println("Exiting the program.")
			os.Exit(0)
		}

		a, err = validateAndParseFloat(input)
		if err != nil {
			fmt.Println("Invalid number. Please enter a valid number.")
			continue
		}

		fmt.Print("Enter second number: ")
		_, err = fmt.Scan(&input)
		if err != nil {
			log.Fatalf("Invalid input for second number: %v", err)
		}

		b, err = validateAndParseFloat(input)
		if err != nil {
			fmt.Println("Invalid number. Please enter a valid number.")
			continue
		}

		operator = getOperatorFromUser()

		bc := calculator.BasicCalculator{}

		result, err := bc.Compute(a, b, operator)
		if err != nil {
			log.Fatalf("Error while computing: %v", err)
		}

		printResult(result)
	}
}

func validateAndParseFloat(input string) (float64, error) {
	input = strings.TrimSpace(input)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number format")
	}
	return value, nil
}

func getOperatorFromUser() string {
	var operator string
	for {
		fmt.Print("Enter operator (+, -, *, /, %, ^): ")
		_, err := fmt.Scan(&operator)
		if err != nil {
			log.Fatalf("Invalid input for operator: %v", err)
		}
		if isValidOperator(operator) {
			return operator
		} else {
			fmt.Println("Invalid operator. Please enter a valid operator (+, -, *, /, %, ^).")
		}
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

func printResult(result float64) {
	if result == float64(int(result)) {
		fmt.Printf("Result: %d\n", int(result))
	} else {
		fmt.Printf("Result: %.4f\n", result)
	}
}
