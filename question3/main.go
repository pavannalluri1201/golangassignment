package main

import "fmt"

func findPosition(arr []int, target int) int {
	for i, num := range arr {
		if num == target {
			return i
		}
	}
	return -1
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	targetNumber := 30

	position := findPosition(numbers, targetNumber)

	fmt.Printf("The position of %d in the array is %d.\n", targetNumber, position)
}
