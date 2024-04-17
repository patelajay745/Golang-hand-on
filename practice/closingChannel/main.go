package main

import "fmt"

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs

			if more {
				fmt.Println("recieved job", j)
			} else {
				fmt.Println("recieved all jobs")
				done <- true
				return
			}

		}
	}()

	for k := range 3 {
		jobs <- k
		fmt.Println("Sent job", k)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("recived more jobs:", ok)

}
