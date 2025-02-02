package main

import (
	"calculator/internal/calculator"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Print("Enter expression (or type 'exit' to quit): ")
		var input string
		_, err := fmt.Scan(&input)

		if err != nil {
			log.Fatalf("Invalid input: %v", err)
		}

		if input == "exit" {
			fmt.Println("Exiting the program.")
			os.Exit(0)
		}

		result, err := calculateExpression(input)
		if err != nil {
			fmt.Println("Invalid expression. Please enter a valid expression.")
			continue
		}

		printResult(result)
	}
}

func calculateExpression(expression string) (float64, error) {
	expression = strings.TrimSpace(expression)

	tokens := tokenize(expression)

	var values []float64
	var operators []string

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if isNumber(token) {
			value, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number format")
			}
			values = append(values, value)
		} else if isOperator(token) {
			for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(token) {
				val2 := values[len(values)-1]
				values = values[:len(values)-1]
				val1 := values[len(values)-1]
				values = values[:len(values)-1]
				operator := operators[len(operators)-1]
				operators = operators[:len(operators)-1]

				result, err := calculator.BasicCalculator{}.Compute(val1, val2, operator)
				if err != nil {
					return 0, fmt.Errorf("error while computing: %v", err)
				}
				values = append(values, result)
			}
			operators = append(operators, token)
		} else if token == "(" {
			operators = append(operators, token)
		} else if token == ")" {
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				val2 := values[len(values)-1]
				values = values[:len(values)-1]
				val1 := values[len(values)-1]
				values = values[:len(values)-1]
				operator := operators[len(operators)-1]
				operators = operators[:len(operators)-1]

				result, err := calculator.BasicCalculator{}.Compute(val1, val2, operator)
				if err != nil {
					return 0, fmt.Errorf("error while computing: %v", err)
				}
				values = append(values, result)
			}
			if len(operators) > 0 && operators[len(operators)-1] == "(" {
				operators = operators[:len(operators)-1]
			} else {
				return 0, fmt.Errorf("mismatched parentheses")
			}
		}
	}

	for len(operators) > 0 {
		val2 := values[len(values)-1]
		values = values[:len(values)-1]
		val1 := values[len(values)-1]
		values = values[:len(values)-1]
		operator := operators[len(operators)-1]
		operators = operators[:len(operators)-1]

		result, err := calculator.BasicCalculator{}.Compute(val1, val2, operator)
		if err != nil {
			return 0, fmt.Errorf("error while computing: %v", err)
		}
		values = append(values, result)
	}

	return values[0], nil
}

func tokenize(expression string) []string {
	re := regexp.MustCompile(`\d+(\.\d+)?|[+\-*/%^()]+`)
	matches := re.FindAllString(expression, -1)
	return matches
}

func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func isOperator(token string) bool {
	operators := "+-*/%^"
	return strings.Contains(operators, token)
}

func precedence(operator string) int {
	switch operator {
	case "^":
		return 3
	case "*", "/":
		return 2
	case "+", "-":
		return 1
	case "%":
		return 2
	default:
		return 0
	}
}

func printResult(result float64) {
	if result == float64(int(result)) {
		fmt.Printf("Result: %d\n", int(result))
	} else {
		fmt.Printf("Result: %.4f\n", result)
	}
}
