package main

import "fmt"

type User struct {
	Name string
	City string
	Car  string
}

func main() {
	fmt.Println("welcome to structs example")
	ajay := User{"Ajay", "London", "Honda"}

	fmt.Println("Details :", ajay)
	fmt.Printf("Another way to show details: %+v\n", ajay)
	fmt.Printf("Another way %v\n", ajay.Name)

	// Modify a field of the struct instance
	ajay.City = "New York"

	// Print modified details of ajay
	fmt.Println("Modified details:", ajay)

	// Initialize a struct instance with zero values for fields
	var emptyUser User

	// Print details of emptyUser
	fmt.Println("Empty User:", emptyUser)

	// Initialize a struct instance using struct literal with selected fields
	robert := User{Name: "Robert"}

	// Print details of robert
	fmt.Println("Robert's details:", robert)

	// Initialize a struct instance using new keyword (returns a pointer to the struct)
	pointerToUser := new(User)

	// Set values of fields using pointerToUser
	pointerToUser.Name = "John"
	pointerToUser.City = "Paris"
	pointerToUser.Car = "BMW"

	// Print details of pointerToUser
	fmt.Println("Pointer to User:", *pointerToUser)
}
