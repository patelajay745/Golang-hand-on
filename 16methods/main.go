package main

import "fmt"

type User struct {
	Name string
	City string
	Car  string
}

func (u User) GetCity() {
	fmt.Println("City is : ", u.City)
}

// Another method associated with the User struct
func (u *User) ChangeCity(newCity string) {
	// Method to change the City field of the User struct
	u.City = newCity
}

func main() {

	fmt.Println("welcome to methos example")
	ajay := User{"Ajay", "London", "Honda"}

	fmt.Println("Details :", ajay)

	ajay.GetCity()

	ajay.ChangeCity("New York")

	ajay.GetCity()

	// Demonstrate methods with pointer receivers
	// Create a pointer to the User struct
	ptrAjay := &ajay
	// Call the ChangeCity method with a pointer receiver
	ptrAjay.ChangeCity("Los Angeles")
	// Print updated details of the User after changing the city using pointer receiver
	fmt.Println("Updated details (using pointer receiver):", ajay)

}
