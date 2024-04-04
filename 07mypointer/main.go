package main

import "fmt"

func main() {

	fmt.Println("welcome to example og pointer")

	// to create just pointer
	// var ptr *int
	// it will occupoy memory location which can hold int value.

	myNumber := 500

	// pointer to refrence the memory location
	var ptr = &myNumber
	fmt.Println("Actual location of variable in menory", ptr)

	*ptr = *ptr / 5
	fmt.Println("modified value of variable: ", *ptr)

}
