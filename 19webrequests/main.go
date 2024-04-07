package main

import (
	"fmt"
	"io"
	"net/http"
)

const Url = "https://www.freecodecamp.org/the-fastest-web-page-on-the-internet"

func main() {
	response, _ := http.Get(Url)

	defer response.Body.Close()
	databyte, _ := io.ReadAll(response.Body)

	fmt.Println("Content:", string(databyte))
}
