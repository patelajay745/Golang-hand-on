package main

import "fmt"

func main() {
	fmt.Println("welcome to if else example")
	numberofLogin := 25
	if numberofLogin > 10 {
		fmt.Println("login is too high")
	} else if numberofLogin < 5 {
		fmt.Println("login is too low")
	} else {
		fmt.Println("login is okay")
	}

	// If statement with a short variable declaration
	if loginCount := 20; loginCount > 15 {
		fmt.Println("Login count is greater than 15")
	} else {
		fmt.Println("Login count is not greater than 15")
	}

	// If statement with multiple conditions
	var userName string = "john"
	var password string = "secret"

	if userName == "john" && password == "secret" {
		fmt.Println("Valid username and password")
	} else if userName == "john" || password == "admin" {
		fmt.Println("Either username or password is correct")
	} else {
		fmt.Println("Invalid username or password")
	}

	// If statement with a condition that evaluates to a boolean value
	if isLoggedIn := true; isLoggedIn {
		fmt.Println("User is logged in")
	}

	// If statement with a condition that evaluates to a boolean value, using a function
	if isEven(10) {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	isOdd(9)

}

// Function to check if a number is even
func isEven(num int) bool {
	return num%2 == 0
}

func isOdd(num int) {
	if num%2 == 0 {
		fmt.Println("Not odd")
	} else {
		fmt.Println("Odd")
	}
}
