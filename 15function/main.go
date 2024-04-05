package main

import (
	"fmt"
)

func main() {
	proresult, promessage := proAdder(10, 65, 89, 4456, 648)
	fmt.Printf("result: %v and message: %v", proresult, promessage)

	// Demonstrate named return values
	namedResult, namedMessage := namedReturnFunction(3, 7)
	fmt.Printf("Named Result: %v and Named Message: %v\n", namedResult, namedMessage)

	// Demonstrate deferred function calls
	deferFunction()

	// Demonstrate anonymous functions
	anonymousFunction()

}

// deferFunction demonstrates deferred function calls
func deferFunction() {
	defer fmt.Println("This will be printed at the end of the function")

	fmt.Println("This will be printed first")
}

func proAdder(value ...int) (int, string) {
	result := 0

	for _, v := range value {
		result += v
	}

	return result, "This is hello from function"
}

// namedReturnFunction demonstrates a function with named return values
func namedReturnFunction(a, b int) (sum int, message string) {
	// Calculate the sum of two integers
	sum = a + b

	// Set the message
	message = "This is a named return function"

	// No explicit return statement needed, as named return values are automatically returned
	return
}

// anonymousFunction demonstrates an anonymous function
func anonymousFunction() {
	// Define and immediately invoke an anonymous function
	func() {
		fmt.Println("This is an anonymous function")
	}()
}
