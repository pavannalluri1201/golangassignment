package main

import (
	"calculator"
	"fmt"
)

func main() {

	a := 10.0
	b := 5.0

	result := calculator.Add(a, b)
	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, result)

	result = calculator.Subtract(a, b)
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, result)

	result = calculator.Multiply(a, b)
	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, result)

	result = calculator.Divide(a, b)
	fmt.Printf("%.2f / %.2f = %.2f\n", a, b, result)
}
