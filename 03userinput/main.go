package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating for my program: ")

	for {
		input, _ := reader.ReadString('\n')

		fmt.Println("Thanks for rating between 1 to 5", input)

		numRating, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil || numRating < 1 || numRating > 5 {
			fmt.Println("Invalid input. Please enter a rating between 1 to 5.")
			continue
		}

		fmt.Println("Added 1 to your rating:", numRating+1)
		break
	}

}
