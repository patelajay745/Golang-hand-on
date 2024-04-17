package main

import "fmt"

func main() {

	message := make(chan string, 1)
	signals := make(chan bool)

	select {
	case msg := <-message:
		fmt.Println("Recieved message", msg)
	default:
		fmt.Println("No message recived")
	}

	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no messsage recieved")
	}

	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
