package main

import (
	"fmt"
	"time"
)

func main() {

	// Create an empty slice to store days
	days := []string{}

	// Iterate over days of the week starting from Sunday (0) to Saturday (6)
	for i := 0; i < 7; i++ {
		// Get the day name for the current iteration
		day := time.Now().AddDate(0, 0, i).Weekday().String()

		// Append the day name to the slice
		days = append(days, day)
	}

	fmt.Println("Slices: ", days)

	// Iterating over slice using a traditional for loop with index
	fmt.Println("Iterating over slice using a traditional for loop with index:")
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// Iterating over slice using a range-based for loop with index
	fmt.Println("Iterating over slice using a range-based for loop with index:")
	for i := range days {
		fmt.Println(days[i])
	}

	// Iterating over slice using a range-based for loop with value only
	fmt.Println("Iterating over slice using a range-based for loop with value only:")
	for _, day := range days {
		fmt.Println(day)
	}

	// Iterating using a for loop with a condition
	fmt.Println("Iterating using a for loop with a condition:")
	rougheValue := 1
	for rougheValue < 10 {
		fmt.Println(rougheValue)
		rougheValue++
	}

	// Iterating using an infinite loop with break statement
	fmt.Println("Iterating using an infinite loop with break statement:")
	counter := 0
	for {
		fmt.Println("Counter:", counter)
		counter++
		if counter >= 5 {
			break // Exit the loop if counter is greater than or equal to 5
		}
	}

	// Iterating using a nested loop
	fmt.Println("Iterating using a nested loop:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("(%d, %d) ", i, j)
		}
		fmt.Println()
	}

}
