package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Print("Enter the limit for Fibonacci numbers: ")
	var input string
	fmt.Scanln(&input)

	limit, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	a, b := 0, 1

	fmt.Print("Fibonacci numbers up to ", limit, ": ", a, " ", b)

	for i := 2; i < limit; i++ {
		c := a + b
		fmt.Print(" ", c)
		a, b = b, c
	}

	fmt.Println()
}
