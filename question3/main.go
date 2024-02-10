package main

import "fmt"

func posistion(arr []int, result int) int {
	for i, num := range arr {
		if num == result {
			return i
		}
	}
	return -1
}

func main() {
	a := []int{5, 6, 7, 8, 9, 10}
	b := 9
	location := posistion(a, b)
	if location == -1 {
		fmt.Println("the given value does not exist in the array")
	} else {
		fmt.Printf("the posistion of %d in the given array is at : %d", b, location)
	}
}
