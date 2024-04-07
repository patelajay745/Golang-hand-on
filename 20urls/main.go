package main

import (
	"fmt"
	"net/url"
)

const MyUrl = "http://ajaypatel.live:5000/fun?query=uarethebest"

func main() {
	fmt.Println("Welcome to example of url")

	result, _ := url.Parse(MyUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	for _, val := range qparams {
		fmt.Println(val)
	}
}
