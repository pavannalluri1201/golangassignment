package main

import "fmt"

func main() {
	startingWeekday := "Wednesday"
	numberOfDays := 5

	// Define weekdays
	weekdays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// Find the index of the starting weekday
	startingDay := 0
	for i, day := range weekdays {
		if day == startingWeekday {
			startingDay = i
			break
		}
	}

	// Calculate the result day
	resultDay := (startingDay + numberOfDays) % 7

	// Print the result
	fmt.Printf("The day after %d days after %s is %s.\n", numberOfDays, startingWeekday, weekdays[resultDay])
}
