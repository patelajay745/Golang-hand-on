package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Welcom to example of file operation")

	content := "this is line to be added"

	file, err := os.Create("./myfile.txt")

	checkNillErr((err))

	length, err := io.WriteString(file, content)
	checkNillErr(err)
	fmt.Println("Lenth:", length)
	defer file.Close()

	readFile("./myfile.txt")
}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNillErr(err)
	fmt.Println("Data: ", string(databyte))
}
