package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("welcome to swithcase example")

	// Create a new random number generator with a custom seed (e.g., current time)
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	// Generate a random number of minutes between 1 and 15
	diceNumber := rng.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2")
	case 3:
		fmt.Println("Dice value is 3")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4 ")
	case 5:
		fmt.Println("Dice value is 5 ")
	case 6:
		fmt.Println("Dice value is 6 and you can roll again")
	default:

		fmt.Println("out of box", diceNumber)
	}

	fmt.Println(" day", time.Now().Weekday())

	// Switch-case statement with multiple expressions in cases
	switch day := time.Now().Weekday(); {
	case day == time.Saturday || day == time.Sunday:
		fmt.Println("It's a weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// Switch-case statement with type switch
	var x interface{} = 42

	switch x.(type) {
	case int:
		fmt.Println("x is an integer")
	case string:
		fmt.Println("x is a string")
	default:
		fmt.Println("x is of unknown type")
	}
}
