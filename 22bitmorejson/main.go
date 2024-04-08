package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to json example")
	// Encodejson()
	DecodeJSON()
}

func Encodejson() {
	myCourses := []course{}

	// Append individual course structs to the slice
	myCourses = append(myCourses, course{"Android Development", 299, "google.com", "abc123", []string{"app-development", "android"}})
	myCourses = append(myCourses, course{"Devops Bootcamp", 599, "google.com", "ac123", nil})
	myCourses = append(myCourses, course{"Golang", 199, "google.com", "abc23", []string{"web-development", "golang"}})

	finalJson, _ := json.MarshalIndent(myCourses, "", "\t")

	fmt.Println(string(finalJson))
}

func DecodeJSON() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	jsonDataFromWeb, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// Decode the JSON data into a slice of maps
	var data []map[string]interface{}
	if err := json.Unmarshal(jsonDataFromWeb, &data); err != nil {
		panic(err)
	}

	// Print the decoded data
	for _, obj := range data {
		for key, value := range obj {
			fmt.Printf("%s: %v\n", key, value)
		}
		//fmt.Println() // Print an empty line between each object
	}

}

// if know the data structure of coming data
// func DecodeJSON() {
// 	var mycourse course
// 	json.Unmarshal(jsonDataFromWeb, &mycourse)
// 	fmt.Printf("%#v\n", mycourse)
// }
