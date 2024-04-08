package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const myurl = "http://localhost:8000/get"
const postUrl = "http://localhost:8000/post"

func main() {

	fmt.Println("welcome to example of requests ")

	//PerformGetRequest(url)
	//PerformPostJsonRequest(postUrl)
	PerformPostFormRequest()
}

func PerformGetRequest(websiteUrl string) {

	response, _ := http.Get(websiteUrl)

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("ContentLength:", response.ContentLength)

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write((content))
	fmt.Println("ByteCount:", byteCount)

	fmt.Println("Data:", responseString.String())

}

func PerformPostJsonRequest(websiteUrl string) {

	requestBody := strings.NewReader(`
			{
				 "Name":"Ajay",
				 "course Name":"Go lang"
				 
			}

	`)

	response, _ := http.Post(websiteUrl, "application/json", requestBody)

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	myurl := "http://localhost:8000/postform"
	data := url.Values{}
	data.Add("Name", "Ajay")
	response, _ := http.PostForm(myurl, data)

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
