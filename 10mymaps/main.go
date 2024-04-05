package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["rb"] = "Ruby"

	fmt.Println(languages["js"])

	//to delete
	delete(languages, "js")

	fmt.Println("list: ", languages)

	//loops
	for key, value := range languages {
		fmt.Printf("%v, %v \n", key, value)
	}

}
