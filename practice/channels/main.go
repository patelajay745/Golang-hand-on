package main

import "fmt"

func main() {

	messages := make(chan string)
	messageWithBuffer := make(chan string, 2)

	messageWithBuffer <- "Buffered"
	messageWithBuffer <- "Channel"

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	fmt.Println(<-messageWithBuffer)
	fmt.Println(<-messageWithBuffer)

}
