package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		urlHasHttp := strings.HasPrefix(url, "http://")

		if !urlHasHttp {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch :reading %s %v \n", url, err)
			os.Exit(1)
		}

		resp.Body.Close()
		fmt.Println("StatusCode:", resp.StatusCode)
		fmt.Printf("%s", b)

	}
}
