package calculator

import (
	"errors"
	"math"
)

type BasicCalculator struct{}

func (bc BasicCalculator) Compute(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	case "%":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return math.Mod(a, b), nil
	case "^":
		return math.Pow(a, b), nil
	default:
		return 0, errors.New("unknown operator")
	}
}
