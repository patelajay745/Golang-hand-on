package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var signals = []string{"test"}
var mut sync.Mutex

func main() {

	websites := []string{
		"https://google.com",
		"https://youtube.com",
		"https://facebook.com",
		"https://amazon.com",
		"https://twitter.com",
		"https://instagram.com",
		"https://reddit.com",
		"https://wikipedia.org",
		"https://yahoo.com",
		"https://netflix.com",
		"https://microsoft.com",
		"https://apple.com",
		"https://linkedin.com",
		"https://ebay.com",
		"https://stackoverflow.com",
		"https://github.com",
		"https://cnn.com",
		"https://bbc.com",
		"https://nytimes.com",
		"https://weather.com",
	}

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println(signals)

}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("problem in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
	}

}
