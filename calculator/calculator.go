package calculator

import (
	"fmt"
)

var result float64

func Calci(c string, a float64, b float64) float64 {
	switch c {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Print("Division Not possible because Divisor can't be Zero")
		} else {
			result = a / b
		}
	}
	return result
}
