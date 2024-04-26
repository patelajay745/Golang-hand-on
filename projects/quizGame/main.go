package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	//flag
	csvfilename := flag.String("csv", "problems.csv", "a csv file in the format of 'questions, answer'")

	timeOut := flag.Int("limit", 10, "the time limit for the quiz in seconds")
	flag.Parse()

	openfile, err := os.Open(*csvfilename)
	if err != nil {
		exit(fmt.Sprint("failed to open th file %s\n", *csvfilename))
	}
	problems, _ := csv.NewReader(openfile).ReadAll()

	var score = 0

	for i, problem := range problems {
		question, answer := problem[0], problem[1]

		// TODO: below code should be run with timer on
		userAnswer := askQuestion(i+1, question, *timeOut)

		//check answer
		if strings.TrimSpace(answer) == userAnswer {
			score++
		}
	}

	fmt.Printf("You Scored %v out of %v\n", score, len(problems))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)

}

func askQuestion(questionNumber int, question string, timeout int) string {
	timer := time.NewTimer(time.Duration(timeout) * time.Second)
	defer timer.Stop()

	fmt.Printf("Problem #%d: %s =", questionNumber, question)

	ch := make(chan string)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		ch <- strings.TrimSpace(text)
	}()

	select {
	case <-timer.C:
		fmt.Println("\nTime's up! Moving to the next question.")
		return ""
	case userAnswer := <-ch:
		return userAnswer
	}

}
