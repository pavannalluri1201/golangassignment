package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// Sample array of integers
	numbers := []int{5, 2, 9, 1, 5, 6}

	// Displaying the input array
	fmt.Println("Input Array:", numbers)

	// Sorting the array using Bubble Sort
	bubbleSort(numbers)

	// Displaying the sorted array
	fmt.Println("Sorted Array:", numbers)
}
