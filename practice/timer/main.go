package main

import (
	"fmt"
	"time"
)

func main() {

	// Simulating a user login
	fmt.Println("User logged in")

	// Start a timer for the session
	sessionTimer := time.NewTimer(5 * time.Second)

	go func() {
		// Wait for the session timer to expire
		<-sessionTimer.C

		// Log the user out after session expires
		fmt.Println("User session expired. Logging out...")
	}()

	// Simulate user activity by resetting the session timer when there's activity
	for i := 0; i < 5; i++ {
		fmt.Println("User activity detected")
		// Reset the session timer
		if !sessionTimer.Stop() {
			<-sessionTimer.C
		}
		sessionTimer.Reset(5 * time.Second)
	}

	fmt.Println("User logged out manually")

}
